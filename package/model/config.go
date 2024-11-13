package model

type GameConfiguration struct {
	// InitialHandSize is the number of cards to draw from the deck before the first turn.
	InitialHandSize int `json:"initialHandSize"`

	// CardsDrawnPerTurn is the number of cards to draw from the deck per turn.
	CardsDrawnPerTurn int `json:"cardsDrawnPerTurn"`

	// OnThePlay is whether the simulation is played on the play or not.
	OnThePlay bool `json:"onThePlay"`
}
