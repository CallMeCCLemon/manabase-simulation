package model

// Card Represents a card in a deck. This can be either a land or a non-land.
type Card struct {
	Name     string   `json:"name"`
	Land     *Land    `json:"land,omitempty"`
	NonLand  *NonLand `json:"nonLand,omitempty"`
	Quantity int      `json:"quantity"`
}

func (c *Card) DeepCopy() *Card {
	return &Card{
		c.Name,
		c.Land,
		c.NonLand,
		c.Quantity,
	}
}

// NewCard Creates a new Card instance. Java-ey but useful for removing some of the boilerplate.
func NewCard(name string, land *Land, nonLand *NonLand, quantity int) *Card {
	return &Card{
		Name:     name,
		Land:     land,
		NonLand:  nonLand,
		Quantity: quantity,
	}
}

type InvalidCard struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}
