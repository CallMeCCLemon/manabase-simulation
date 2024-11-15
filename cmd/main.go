package main

import (
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"manabase-simulation/package/model"
	"manabase-simulation/package/reader"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type GameConfiguration struct {
	InitialHandSize   int  `json:"initialHandSize"`
	CardsDrawnPerTurn int  `json:"cardsDrawnPerTurn"`
	OnThePlay         bool `json:"onThePlay"`
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	deck, _ := reader.ReadJSONFile[model.DeckList]("./fixtures/lotus-field-deck.json")
	logger := CreateLogger()
	logger.Info(deck.ToString())
	gameConfig, _ := reader.ReadJSONFile[GameConfiguration]("./fixtures/default-game-config.json")
	objective := model.TestObjective{
		TargetTurn: 3,
		ManaCosts: []model.ManaCost{
			{
				ColorRequirements: []model.ManaColor{model.White, model.White},
				GenericCost:       1,
			},
		},
	}
	now := time.Now()

	successCount := 0
	iterations := 10000

	c := make(chan bool, 100)
	wg := new(sync.WaitGroup)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go start(deck, gameConfig, objective, c, wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for result := range c {
		if result {
			successCount++
		}
	}

	// Capture results to be consumes.
	logger.Info(fmt.Sprintf("Success count: %d", successCount))
	logger.Info(fmt.Sprintf("Success Rate: %f", float32(successCount)/float32(iterations)*100.0))
	logger.Info(fmt.Sprintf("Time taken: %s", time.Since(now)))
}

func start(deckList model.DeckList, gameConfiguration GameConfiguration, objective model.TestObjective, c chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- SimulateDeck(deckList, gameConfiguration, objective)
}

// CreateLogger Creates a new logger with the default configuration.
func CreateLogger() *zap.Logger {
	// Create a custom logger configuration
	config := zap.NewProductionConfig()

	// Set the output to stdout
	config.OutputPaths = []string{"stdout"}

	// Set the error output to stderr
	config.ErrorOutputPaths = []string{"stderr"}

	// Configure the encoder to use a human-readable format (good for testing)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Optional: Colorize log levels

	// Build the logger
	logger, err := config.Build(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		// Redirect the core output to GinkgoWriter
		return zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			zapcore.AddSync(ginkgo.GinkgoWriter), // Send logs to GinkgoWriter
			zapcore.InfoLevel,
		)
	}))
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	// Use the logger
	return logger
}

// SimulateDeck Simulates a deck against a given objective with the provided configuration.
func SimulateDeck(deckList model.DeckList, gameConfiguration GameConfiguration, objective model.TestObjective) bool {
	//logger := CreateLogger()
	//logger.Debug("Starting deck simulation")

	// Generate Randomized Deck
	deck := deckList.GenerateDeck()
	hand := model.NewDeck()
	board := model.NewBoardState()

	// TODO: Add validations like Validate deck is >= 60 cards

	// Draw Initial Hand
	for range gameConfiguration.InitialHandSize {
		hand = deck.DrawCard(hand)
	}

	// For turnNumber to target turn
	for turnNumber := range objective.TargetTurn {
		// If turnNumber = 1 and on the play, skip draw
		if turnNumber == 0 && gameConfiguration.OnThePlay {
			// Skip your draw
			//logger.Debug("Playing first, skipping draw")
		} else {
			hand = deck.DrawCard(hand)
		}

		hand = board.PlayLand(hand, objective, turnNumber+1)
	}

	// Compute if target is met (possibly using backtracking?)
	// Computation can start with the most restrictive lands by sorting based on number of colors it taps for.
	isMet, _ := board.ValidateTestObjective(objective)
	return isMet
}
