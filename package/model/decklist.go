package model

import (
	"encoding/json"
	"slices"
)

// DeckList Represents a playable deck of cards. Quantities are not guaranteed to be 1. This is what a user should be uploading for simplicity.
type DeckList struct {
	Lands    []Land    `json:"lands"`
	NonLands []NonLand `json:"nonLands"`
}

func (d *DeckList) ToString() string {
	jsonPayload, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(jsonPayload)
}

// GetTotalCardCount Returns the total number of cards in the deck.
func (d *DeckList) GetTotalCardCount() int {
	count := 0
	for _, l := range d.Lands {
		count += l.Quantity
	}

	for _, n := range d.NonLands {
		count += n.Quantity
	}

	return count
}

// GenerateDeck Creates a Deck instance from a DeckList.
func (d *DeckList) GenerateDeck() Deck {
	deck := NewDeck()

	for _, l := range d.Lands {
		quantity := l.Quantity
		l.Quantity = 1
		for range quantity {
			deck.Cards = append(deck.Cards, *NewCard(&l, nil))
		}
	}

	for _, n := range d.NonLands {
		quantity := n.Quantity
		n.Quantity = 1
		for range quantity {
			deck.Cards = append(deck.Cards, *NewCard(nil, &n))
		}
	}

	deck.Shuffle()
	return deck
}

// Land Represents a Land type of card which can produce mana.
type Land struct {
	// Name the name of the land.
	Name string `json:"name"`

	// Colors is the list of colors which can be produced by the land.
	Colors []ManaColor `json:"colors"`

	// EntersTapped is whether the land enters tapped.
	EntersTapped bool `json:"entersTapped"`

	// ActivationCost is the list of colors which must be used to activate the land.
	ActivationCost []string `json:"activationCost"`

	// Quantity is the number of copies of this card in a deck.
	Quantity int `json:"quantity"`
}

// Equals Checks if two lands are equal.
func (l *Land) Equals(land Land) bool {
	// TODO: Include other fields.
	if l.EntersTapped != land.EntersTapped {
		return false
	}
	if l.Name != land.Name {
		return false
	}
	if len(l.Colors) != len(land.Colors) {
		return false
	}
	for _, color := range l.Colors {
		if !slices.Contains(land.Colors, color) {
			return false
		}
	}

	return true
}

// NonLand Represents a Non-Land type of card is will need mana to be cast.
type NonLand struct {
	// Name the name of the land.
	Name string `json:"name"`

	// CastingCost the mana which is required to cast the given spell.
	CastingCost []ManaCost `json:"castingCost"`

	// Quantity is the number of copies of this card in a deck.
	Quantity int `json:"quantity"`
}
