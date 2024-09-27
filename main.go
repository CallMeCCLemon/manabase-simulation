package main

import (
	"encoding/json"
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"math/rand/v2"
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

type TestObjective struct {
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
	Land    *Land
	NonLand *NonLand
}

func NewCard(land *Land, nonLand *NonLand) *Card {
	return &Card{
		Land:    land,
		NonLand: nonLand,
	}
}

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	return Deck{
		Cards: []Card{},
	}
}

func (d *Deck) Shuffle() {
	shuffledCards := make([]Card, len(d.Cards))
	perm := rand.Perm(len(d.Cards))

	for i, v := range perm {
		shuffledCards[v] = d.Cards[i]
	}
	d.Cards = shuffledCards
}

// DeepCopy Copies a Deck to a new obj.
func (d *Deck) DeepCopy() Deck {
	newDeck := NewDeck()
	for _, card := range d.Cards {
		newDeck.Cards = append(newDeck.Cards, card)
	}

	return newDeck
}

func (d *Deck) DrawCard(hand Deck) (updatedHand Deck) {
	hand.Cards = append(hand.Cards, d.Cards[0])
	d.Cards = d.Cards[1:]

	return hand
}

type BoardState struct {
	Lands []Land `json:"lands"`
}

func NewBoardState() BoardState {
	return BoardState{
		Lands: make([]Land, 0),
	}
}

// PlayLand plays the best land from the hand based on the turn and target condition
func (b *BoardState) PlayLand(hand Deck, objective TestObjective, turn int) (updatedHand Deck) {
	// Play a Land. If target turn, prioritize untapped. If not, prioritize tapped.
	//  Prioritize lands which generate colors in mana costs.
	// Update Hand.
	// Update Board State

	return hand
}

// ValidateTestObjective validates whether or not the TestObjective has been met
func (b *BoardState) ValidateTestObjective(objective TestObjective) bool {
	return true
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
	return board.ValidateTestObjective(objective)
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
