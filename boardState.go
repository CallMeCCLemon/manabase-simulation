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
	lands := make([]Land, 0)
	nonLands := make([]Card, 0)
	newHand := NewDeck()
	for _, c := range hand.Cards {
		if c.Land != nil {
			lands = append(lands, *c.Land)
		} else {
			nonLands = append(nonLands, c)
		}
	}

	prioritizeUntapped := false
	if turn == objective.TargetTurn {
		prioritizeUntapped = true
	}

	_, combos := b.ValidateTestObjective(objective)

	i, l := b.selectLand(lands, combos, prioritizeUntapped)
	if i < 0 {
		// This implies no lands in hand. Womp womp.
		return hand
	}

	lands = slices.Delete(lands, i, i+1)
	newHand.Cards = nonLands
	for _, land := range lands {
		newHand.Cards = append(newHand.Cards, *NewCard(&land, nil))
	}
	b.Lands = append(b.Lands, l)
	return newHand
}

// selectLand selects a land to play based on the lands score. Will take into account if it should prioritize untapped lands.
func (b *BoardState) selectLand(lands []Land, costOptions []ManaCost, prioritizeUntapped bool) (int, Land) {
	// try to select an untapped land
	var remainingLands []Land
	if prioritizeUntapped {
		for _, l := range lands {
			if !l.EntersTapped {
				remainingLands = append(remainingLands, l)
			}
		}
	}

	if len(remainingLands) == 0 {
		// otherwise accept any land choice.
		remainingLands = lands
	}

	slices.SortFunc(remainingLands, func(a Land, b Land) int {
		scoreA := scoreLand(a, costOptions)
		scoreB := scoreLand(b, costOptions)
		if scoreA < scoreB {
			return -1
		} else if scoreA == scoreB {
			return 0
		} else {
			return 1
		}
	})
	targetLand := remainingLands[len(remainingLands)-1]
	for idx, l := range lands {
		if l.Equals(targetLand) {
			return idx, targetLand
		}
	}
	return -1, targetLand
}

// scoreLand scores a land based on how well it can resolve existing costs.
func scoreLand(l Land, costOptions []ManaCost) int {
	// TODO: Consider using a weighted avg per-costOption which computes the number of colors provided / total # colors for the option
	score := 0
	for _, cost := range costOptions {
		evaluatedColors := make(map[ManaColor]bool)
		for _, color := range cost.ColorRequirements {
			if _, ok := evaluatedColors[color]; ok {
				continue
			}

			if slices.Contains(l.Colors, color) {
				score += 1
				evaluatedColors[color] = true
			}
		}
	}

	// Heavily weight lands which can produce pips for target, but also give a higher score to lands which produce multiple colors.
	score = 10*score + len(l.Colors)

	return score
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
			if cost.GetRemainingCost() == 0 {
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
			key := cost.ToString()
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
		if cost.GetRemainingCost() == 0 {
			isObjectiveMet = true
			upcomingManaCosts = nil
			return true, nil
		}
		if cost.GetRemainingCost() < minSize {
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

// indexOf finds the index of a specific value in a slice. If not found, returns -1.
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
