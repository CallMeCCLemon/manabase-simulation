package model

// Card Represents a card in a deck. This can be either a land or a non-land.
type Card struct {
	Name    string   `json:"name"`
	Land    *Land    `json:"land,omitempty"`
	NonLand *NonLand `json:"nonLand,omitempty"`
}

// NewCard Creates a new Card instance. Java-ey but useful for removing some of the boilerplate.
func NewCard(land *Land, nonLand *NonLand) *Card {
	return &Card{
		Land:    land,
		NonLand: nonLand,
	}
}
