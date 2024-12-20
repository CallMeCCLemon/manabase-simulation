package validation

import (
	"fmt"
	"gorm.io/driver/postgres"
	"manabase-simulation/package/model"
	"regexp"
	"strconv"
	"strings"
)

const (
	MissingCardReason = "card not found in standard card DB"
)

// Parser Represents a utility to parse a string into a model.DeckList
type Parser interface {
	// Parse main function to parse a string into a model.DeckList.
	Parse(string) (*model.DeckList, []model.InvalidCard, error)

	// SeparateDeckAndSideboard parses out the deck and sideboard from a list of strings (assuming MTGA Decklist format).
	SeparateDeckAndSideboard(deckList []string) ([]string, []string, error)
}

// DefaultParser Represents a default implementation of the Parser interface.
type DefaultParser struct {
	CardDbAccessor CardDbAccessor
}

var _ Parser = &DefaultParser{}

func NewDefaultParser(psqlConfig postgres.Config) *DefaultParser {
	db, err := NewPsqlDbAccessor(psqlConfig)
	if err != nil {
		panic(err)
	}
	return &DefaultParser{
		CardDbAccessor: db,
	}
}

func (d *DefaultParser) Parse(s string) (*model.DeckList, []model.InvalidCard, error) {
	// Split all lines into a slice.
	lines := strings.Split(s, "\n")

	// For each line, separate the deck from sideboard.
	deckLines, _, err := d.SeparateDeckAndSideboard(lines)
	if err != nil {
		return nil, nil, err
	}

	// Get a slice of all the maindeck cards and their quantities as a map.
	deckCardsAndQuantities, err := d.getNamesAndQuantities(deckLines)
	if err != nil {
		return nil, nil, err
	}

	// Use the keys of names to get the cards from the card database.
	deckCardNames := make([]string, len(deckCardsAndQuantities))
	i := 0
	for name := range deckCardsAndQuantities {
		deckCardNames[i] = name
		i++
	}
	deckCards, err := d.CardDbAccessor.GetCards(deckCardNames)
	if err != nil {
		return nil, nil, err
	}

	invalidCards := make([]model.InvalidCard, 0)

	// Update the cards in the deck with their quantities.
	for name := range deckCardsAndQuantities {
		card := deckCards[name]
		if card == nil {
			invalidCards = append(invalidCards, model.InvalidCard{
				Name:   name,
				Reason: MissingCardReason,
			})
		} else {
			card.Quantity = deckCardsAndQuantities[name]
		}
	}

	var cards []model.Card
	for _, card := range deckCards {
		if card != nil {
			cards = append(cards, *card)
		}
	}

	return &model.DeckList{
		Cards: cards,
	}, invalidCards, nil
}

func (d *DefaultParser) SeparateDeckAndSideboard(deckList []string) ([]string, []string, error) {
	isDeck := false
	isSideboard := false

	var deckLines []string
	var sideboardLines []string

	for _, line := range deckList {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Sideboard") {
			isDeck = false
			isSideboard = true
			continue
		}

		if strings.HasPrefix(line, "Deck") {
			isDeck = true
			isSideboard = false
			continue
		}

		if isDeck {
			deckLines = append(deckLines, line)
		}

		if isSideboard {
			sideboardLines = append(sideboardLines, line)
		}
	}

	return deckLines, sideboardLines, nil
}

func (d *DefaultParser) getNamesAndQuantities(deckList []string) (map[string]int, error) {
	names := make(map[string]int, len(deckList))

	re := regexp.MustCompile(`^(\d+)?\s*(.*)`)

	for _, line := range deckList {
		// Find the matches
		matches := re.FindStringSubmatch(line)

		// Initialize number with default value 1
		number := 1

		if len(matches) < 3 {
			return nil, fmt.Errorf("invalid deck line: %s", line)
		}
		// If a number was captured, parse it
		var text string
		if matches[1] != "" {
			n, err := strconv.Atoi(matches[1])
			if err == nil {
				number = n
			}
		}
		text = strings.TrimSpace(matches[2])
		// The card name is the second capture group, trimmed of whitespace

		names[text] = number
	}

	return names, nil
}
