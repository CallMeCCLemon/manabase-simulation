package main

// TestObjective is the objective to be met by the deck during the simulation. Currently, this is only for a single card, but this could later be extended to be for all non-land cards in the deck.
type TestObjective struct {
	TargetTurn int        `json:"targetTurn"`
	ManaCosts  []ManaCost `json:"manaCosts"`
}
