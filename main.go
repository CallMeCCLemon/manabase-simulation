package main

import (
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sort"
	"sync"
	"time"
)

type GameConfiguration struct {
	InitialHandSize   int  `json:"initialHandSize"`
	CardsDrawnPerTurn int  `json:"cardsDrawnPerTurn"`
	OnThePlay         bool `json:"onThePlay"`
}

// SortLandsByRestrictiveness Sorts a list of lands by the number of colors they produce.
func SortLandsByRestrictiveness(lands []Land) []Land {
	sort.Slice(lands, func(i, j int) bool {
		return len(lands[i].Colors) < len(lands[j].Colors)
	})

	return lands
}

// ManaColor Represents a color of mana in the game.
type ManaColor string

const (
	// white Represents the white color of mana.
	white ManaColor = "white"

	// blue Represents the blue color of mana.
	blue ManaColor = "blue"

	// black Represents the black color of mana.
	black ManaColor = "black"

	// red Represents the red color of mana.
	red ManaColor = "red"

	// green Represents the green color of mana.
	green ManaColor = "green"

	// colorless Represents the colorless mana.
	colorless ManaColor = "colorless"

	// whatever Represents any color of mana. This is used primarily for wildcard mana producers, but I'm unsure if this is really necessary.
	whatever ManaColor = "whatever"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	deck, _ := ReadJSONFile[DeckList]("./fixtures/lotus-field-deck.json")
	logger := CreateLogger()
	logger.Info(deck.ToString())
	gameConfig, _ := ReadJSONFile[GameConfiguration]("./fixtures/default-game-config.json")
	objective := TestObjective{
		TargetTurn: 3,
		ManaCosts: []ManaCost{
			{
				ColorRequirements: []ManaColor{white, white},
				GenericCost:       1,
			},
		},
	}
	now := time.Now()

	successCount := 0
	iterations := 10000000

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

func start(deckList DeckList, gameConfiguration GameConfiguration, objective TestObjective, c chan bool, wg *sync.WaitGroup) {
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

func ReadJSONFile[T any](filename string) (T, error) {
	var data T
	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return data, err
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// SimulateDeck Simulates a deck against a given objective with the provided configuration.
func SimulateDeck(deckList DeckList, gameConfiguration GameConfiguration, objective TestObjective) bool {
	//logger := CreateLogger()
	//logger.Debug("Starting deck simulation")

	// Generate Randomized Deck
	deck := deckList.GenerateDeck()
	hand := NewDeck()
	board := NewBoardState()

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

// indexOf finds the index of a specific value in a slice. If not found, returns -1.
func indexOf[T comparable](slice []T, value T) int {
	for index, v := range slice {
		if v == value {
			return index
		}
	}
	return -1
}
