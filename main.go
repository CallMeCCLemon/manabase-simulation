package main

import (
	"encoding/json"
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"os"
)

type DeckList struct {
	Lands    []Land    `json:"lands"`
	NonLands []NonLand `json:"nonLands"`
}

func NewDeckList() *DeckList {
	return &DeckList{
		Lands:    []Land{},
		NonLands: []NonLand{},
	}
}

func (d *DeckList) toString() string {
	jsonPayload, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(jsonPayload)
}

func (d *DeckList) getTotalCardCount() int {
	count := 0
	for _, l := range d.Lands {
		count += l.Quantity
	}

	for _, n := range d.NonLands {
		count += n.Quantity
	}

	return count
}

type Land struct {
	Name           string   `json:"name"`
	Colors         []string `json:"colors"`
	EntersTapped   bool     `json:"entersTapped"`
	ActivationCost []string `json:"activationCost"`
	Quantity       int      `json:"quantity"`
}

type NonLand struct {
	Name        string   `json:"name"`
	CastingCost []string `json:"castingCost"`
	Quantity    int      `json:"quantity"`
}

type TestCondition struct {
	TargetTurn int        `json:"targetTurn"`
	ManaCosts  []ManaCost `json:"manaCosts"`
}

type ManaCost struct {
	ColorRequirements []string `json:"colorRequirements"`
	GenericCost       int      `json:"genericCost"`
}

type GameConfiguration struct {
	InitialHandSize   int  `json:"initialHandSize"`
	CardsDrawnPerTurn int  `json:"cardsDrawnPerTurn"`
	OnThePlay         bool `json:"onThePlay"`
}

type Card struct {
	land    *Land
	nonLand *NonLand
}

type Deck struct {
	cards []Card
}

func NewDeck() Deck {
	return Deck{
		cards: []Card{},
	}
}

type BoardState struct {
	Lands []Land `json:"lands"`
}

const (
	white     = "white"
	blue      = "blue"
	black     = "black"
	red       = "red"
	green     = "green"
	colorless = "colorless"
	whatever  = "whatever"
)

func main() {
	println("Hello World")
	deck, _ := ReadDeckJSON("./sample_deck.json")
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

// ReadDeckJSON Function to read JSON file into a struct
func ReadDeckJSON(filename string) (DeckList, error) {
	var deck DeckList

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return deck, err
	}
	defer file.Close()

	// Read the file contents
	bytes, err := ioutil.ReadAll(file)
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
	bytes, err := ioutil.ReadAll(file)
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

func SimulateDeck(deckList DeckList, gameConfiguration GameConfiguration) {
	logger := createLogger()
	logger.Info("Starting deck simulation", zap.String("deck", deckList.toString()))

	// Generate Randomized Deck
	deck := GenerateDeck(deckList)

	hand := NewDeck()

	// Draw Initial Hand
	// For i to initial hand size
	// Draw cards

	// For i to target turn
	// If i = 1 and on the play, skip draw
	if !gameConfiguration.OnThePlay {
		// Draw another card
		deck, hand = DrawCard(deck, hand)
	}
	// Play a land. If target turn, prioritize untapped. If not, prioritize tapped.
	//  Prioritize lands which generate colors in mana costs.
	// Update Hand.
	// Update Board State

	// Repeat until target turn

	// Compute if target is met (possibly using backtracking?)
	// Computation can start with the most restrictive lands by sorting based on number of colors it taps for.
}

func GenerateDeck(list DeckList) Deck {
	deck := NewDeck()

	return deck
}

func ShuffleDeck(deck Deck) Deck {
	return deck
}

func DrawCard(deck Deck, hand Deck) (updatedDeck Deck, updatedHand Deck) {
	return Deck{}, Deck{}
}
