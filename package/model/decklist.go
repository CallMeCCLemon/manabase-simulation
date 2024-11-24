package model

import (
	"encoding/json"
	"errors"
	"manabase-simulation/package/reader"
	"slices"
)

// DeckList Represents a playable deck of cards. Quantities are not guaranteed to be 1. This is what a user should be uploading for simplicity.
type DeckList struct {
	Lands    []Land    `json:"lands"`
	NonLands []NonLand `json:"nonLands"`
}

// ToString returns a string representation of the deck.
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

	// Types is a list of basic land types which the current land is.
	Types []LandType `json:"types,omitempty"`

	// Colors is the list of colors which can be produced by the land.
	Colors []ManaColor `json:"colors,omitempty"`

	// EntersTapped is whether the land enters tapped.
	EntersTapped bool `json:"entersTapped"`

	// ActivationCost is the cost which must be paid to activate the land.
	ActivationCost *ActivationCost `json:"activationCost,omitempty"`

	// UntappedCondition is the condition which must be met to enter untapped.
	UntappedCondition *UntappedCondition `json:"untappedCondition,omitempty"`

	// Quantity is the number of copies of this card in a deck.
	Quantity int `json:"quantity"`
}

// ActivationCost Represents the cost which must be paid to activate the land.
type ActivationCost struct {
	Life *int      `json:"life,omitempty"`
	Mana *ManaCost `json:"manaCost,omitempty"`
}

// UntappedCondition Represents a condition which must be met to enter untapped.
type UntappedCondition struct {
	Type ConditionType `json:"type"`
	Data *string       `json:"data,omitempty"`
}

// LandType is an enum of basic land types.
type LandType string

const (
	// Plains is a basic land type which produces white mana.
	Plains LandType = "Plains"

	// Mountain is a basic land type which produces red mana.
	Mountain LandType = "Mountain"

	// Forest is a basic land type which produces green mana.
	Forest LandType = "Forest"

	// Island is a basic land type which produces blue mana.
	Island LandType = "Island"

	// Swamp is a basic land type which produces black mana.
	Swamp LandType = "Swamp"
)

// ConditionType Represents the type of condition for a land to enter untapped.
type ConditionType string

const (
	// ShockLand is a condition where a land enters untapped if 2 life is paid.
	ShockLand ConditionType = "ShockLand"

	// FastLand is a condition where a land enters untapped if the total number of lands before playing is 2 or less.
	FastLand ConditionType = "FastLand"

	// CheckLand is a condition where a land enters untapped if there is a land of the specified type on the board.
	CheckLand ConditionType = "CheckLand"
)

// CheckLandData is supplemental data provided with a CheckLand condition.
type CheckLandData []LandType

func (c *CheckLandData) ToString() (string, error) {
	jsonPayload, err := reader.WriteJSONString(*c)
	if err != nil {
		return "", err
	}

	return jsonPayload, nil
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

// PayUntappedCost Pays the cost of the land to enter untapped.
func (l *Land) PayUntappedCost(b *BoardState) error {
	if l.EntersTapped == false {
		return nil
	}

	if l.UntappedCondition == nil {
		return errors.New("untapped condition not found")
	}

	// Switch for all the untapped conditions.
	switch l.UntappedCondition.Type {
	case ShockLand:
		if b.Life > 2 {
			b.Life -= 2
			return nil
		} else {
			return errors.New("not enough life to enter untapped")
		}
	case FastLand:
		if l.CanEnterUntapped(*b) {
			return nil
		} else {
			return errors.New("too many lands to enter untapped")
		}
	case CheckLand:
		if l.CanEnterUntapped(*b) {
			return nil
		}

		return errors.New("no lands of the right type found")
	default:
		return errors.New("unknown untapped condition")
	}
}

// CanEnterUntapped checks if the land can enter untapped based on the BoardState and UntappedCondition.
func (l *Land) CanEnterUntapped(b BoardState) bool {
	if l.EntersTapped == false {
		return true
	}

	if l.UntappedCondition == nil {
		return false
	}

	switch l.UntappedCondition.Type {
	case ShockLand:
		if b.Life > 2 {
			return true
		} else {
			return false
		}
	case FastLand:
		return len(b.Lands) <= 2
	case CheckLand:
		if l.UntappedCondition.Data == nil {
			return false
		}
		c, err := reader.ReadJSONString[CheckLandData](*l.UntappedCondition.Data)
		if err != nil {
			return false
		}

		for _, lands := range b.Lands {
			for _, t := range c {
				if slices.Contains(lands.Types, t) {
					return true
				}
			}
		}
		// No lands of the right type found.
		return false
	default:
		return false
	}
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
