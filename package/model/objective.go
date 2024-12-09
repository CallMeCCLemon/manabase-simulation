package model

import "fmt"

// TestObjective is the objective to be met by the deck during the simulation. Currently, this is only for a single card, but this could later be extended to be for all non-land cards in the deck.
type TestObjective struct {
	// TargetTurn is the number of turns at which the objective should be met.
	TargetTurn int `json:"targetTurn"`

	// ManaCosts is the list of mana costs which must be satisfied to meet the objective.
	ManaCosts []ManaCost `json:"manaCosts"`
}

func (o *TestObjective) ToString() string {
	return fmt.Sprintf("TargetTurn: %d, ManaCosts: %v", o.TargetTurn, o.ManaCosts)
}
