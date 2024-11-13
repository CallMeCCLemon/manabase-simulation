package model

// Card Represents a card in a deck. This can be either a land or a non-land.
type Card struct {
	Land    *Land
	NonLand *NonLand
}

// NewCard Creates a new Card instance. Java-ey but useful for removing some of the boilerplate.
func NewCard(land *Land, nonLand *NonLand) *Card {
	return &Card{
		Land:    land,
		NonLand: nonLand,
	}
}
