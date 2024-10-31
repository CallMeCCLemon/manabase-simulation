package main

import (
	"slices"
)

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
	lands := make([]Card, 0)
	nonLands := make([]Card, 0)
	newHand := NewDeck()
	for _, c := range hand.Cards {
		if c.Land != nil {
			lands = append(lands, c)
		} else {
			nonLands = append(nonLands, c)
		}
	}

	//_ := b.GetManaCombinations()

	//  Prioritize lands which generate colors in mana costs.
	//prioritizeUntapped := false
	//if turn == objective.TargetTurn {
	//	prioritizeUntapped = true
	//}
	//
	//for idx, l := range lands {
	//	combinations := b.GetManaCombinations()
	//
	//}

	// Update Hand.
	// Update Board State

	return newHand
}

// ValidateTestObjective validates if the TestObjective has been met
func (b *BoardState) ValidateTestObjective(objective TestObjective) (bool, []ManaCost) {
	// TODO: Evaluate how we can do this for multiple costs.

	// Sort lands by most restrictive production where the first have most restricted colors
	sortedLands := SortLandsByRestrictiveness(b.Lands)

	manaCosts := make([]ManaCost, 0)
	upcomingManaCosts := []ManaCost{objective.ManaCosts[0]}
	for _, l := range sortedLands {
		manaCosts = upcomingManaCosts
		upcomingManaCosts = make([]ManaCost, 0)
		for _, cost := range manaCosts {
			if len(cost.ColorRequirements) == 0 && cost.GenericCost == 0 {
				return true, nil
			}
			for _, color := range l.Colors {
				// Use the land to remove a color if possible
				if slices.Contains(cost.ColorRequirements, color) {
					idx := indexOf(cost.ColorRequirements, color)
					tmpManaCost := cost.DeepCopy()
					tmpManaCost.ColorRequirements = slices.Delete(tmpManaCost.ColorRequirements, idx, idx+1)
					upcomingManaCosts = append(upcomingManaCosts, tmpManaCost)
				} else {
					// Fallback to consuming a generic cost if necessary
					if cost.GenericCost > 0 {
						tmpManaCost := ManaCost{
							ColorRequirements: cost.ColorRequirements,
							GenericCost:       cost.GenericCost - 1,
						}
						upcomingManaCosts = append(upcomingManaCosts, tmpManaCost)
					} else {
						upcomingManaCosts = append(upcomingManaCosts, cost)
					}
				}
			}
		}
		// Remove Dupes
		dedupedManaCosts := make(map[string]ManaCost)
		for _, cost := range upcomingManaCosts {
			key := cost.toString()
			if _, ok := dedupedManaCosts[key]; !ok {
				dedupedManaCosts[key] = cost
			}
		}
		upcomingManaCosts = make([]ManaCost, 0)
		for _, value := range dedupedManaCosts {
			upcomingManaCosts = append(upcomingManaCosts, value)
		}
	}

	// Check the remaining mana costs for a completed one
	// Prune all decisions which have costs > minCostSize
	minSize := 100
	isObjectiveMet := false
	for _, cost := range upcomingManaCosts {
		if cost.GenericCost == 0 && len(cost.ColorRequirements) == 0 {
			isObjectiveMet = true
			upcomingManaCosts = nil
			return true, nil
		}
		if len(cost.ColorRequirements)+cost.GenericCost < minSize {
			minSize = len(cost.ColorRequirements) + cost.GenericCost
		}
	}

	remainingManaCosts := make([]ManaCost, 0)
	for _, cost := range upcomingManaCosts {
		if len(cost.ColorRequirements)+cost.GenericCost == minSize {
			remainingManaCosts = append(remainingManaCosts, cost)
		}
	}

	return isObjectiveMet, remainingManaCosts
}

func indexOf[T comparable](slice []T, value T) int {
	for index, v := range slice {
		if v == value {
			return index
		}
	}
	return -1
}

// GetManaCombinations returns all the possible mana combinations.
func (b *BoardState) GetManaCombinations() [][]ManaColor {
	manaCombos := make([][]ManaColor, 0)
	// Use all the lands on the board to build the combos
	for _, l := range b.Lands {
		// Use each color the land could produce to create possible combos
		for _, color := range l.Colors {
			tmpCombos := make([][]ManaColor, 0)
			if len(manaCombos) == 0 {
				// Create 1 color combos
				manaCombos = append(manaCombos, []ManaColor{ManaColor(color)})
			} else {
				// Create a new combo with the land color value
				for _, combo := range manaCombos {
					newCombo := append(combo, color)
					slices.Sort(newCombo)
					tmpCombos = append(tmpCombos, newCombo)
				}
			}
			// Overwrite the combos with the update combo lists.
			manaCombos = tmpCombos
		}
	}

	return manaCombos
}
