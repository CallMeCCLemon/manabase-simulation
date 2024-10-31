package main

import (
	"encoding/json"
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"sort"
)

type TestObjective struct {
	TargetTurn int        `json:"targetTurn"`
	ManaCosts  []ManaCost `json:"manaCosts"`
}

type GameConfiguration struct {
	InitialHandSize   int  `json:"initialHandSize"`
	CardsDrawnPerTurn int  `json:"cardsDrawnPerTurn"`
	OnThePlay         bool `json:"onThePlay"`
}

func SortLandsByRestrictiveness(lands []Land) []Land {
	sort.Slice(lands, func(i, j int) bool {
		return len(lands[i].Colors) < len(lands[j].Colors)
	})

	return lands
}

type ManaColor string

const (
	white     ManaColor = "white"
	blue      ManaColor = "blue"
	black     ManaColor = "black"
	red       ManaColor = "red"
	green     ManaColor = "green"
	colorless ManaColor = "colorless"
	whatever  ManaColor = "whatever"
)

func main() {
	println("Hello World")
	deck, _ := ReadDeckListJSON("./sample_deck.json")
	createLogger().Info(deck.toString())
}

func createLogger() *zap.Logger {
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
			zapcore.DebugLevel,
		)
	}))
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	// Use the logger
	logger.Info("This is an info message")
	return logger
}

// ReadDeckListJSON Function to read JSON file into a struct
func ReadDeckListJSON(filename string) (DeckList, error) {
	var deck DeckList

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return deck, err
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return deck, err
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(bytes, &deck)
	if err != nil {
		return deck, err
	}

	return deck, nil
}

// ReadGameConfigJSON Reads a game config JSON file.
func ReadGameConfigJSON(filename string) (GameConfiguration, error) {
	var gameConfig GameConfiguration

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return gameConfig, err
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return gameConfig, err
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(bytes, &gameConfig)
	if err != nil {
		return gameConfig, err
	}

	return gameConfig, nil
}

func SimulateDeck(deckList DeckList, gameConfiguration GameConfiguration, objective TestObjective) bool {
	logger := createLogger()
	logger.Info("Starting deck simulation", zap.String("deck", deckList.toString()))

	// Generate Randomized Deck
	deck := GenerateDeck(deckList)
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
		if turnNumber == 1 && gameConfiguration.OnThePlay {
			// Skip your draw
			logger.Info("Playing first, skipping draw")
		} else {
			hand = deck.DrawCard(hand)
		}

		board.PlayLand(hand, objective, turnNumber)
	}

	// Compute if target is met (possibly using backtracking?)
	// Computation can start with the most restrictive lands by sorting based on number of colors it taps for.
	isMet, _ := board.ValidateTestObjective(objective)
	return isMet
}

// GenerateDeck Creates a Deck instance from a DeckList.
func GenerateDeck(list DeckList) Deck {
	deck := NewDeck()

	for _, l := range list.Lands {
		quantity := l.Quantity
		l.Quantity = 1
		for range quantity {
			deck.Cards = append(deck.Cards, *NewCard(&l, nil))
		}
	}

	for _, n := range list.NonLands {
		quantity := n.Quantity
		n.Quantity = 1
		for range quantity {
			deck.Cards = append(deck.Cards, *NewCard(nil, &n))
		}
	}

	deck.Shuffle()
	return deck
}
