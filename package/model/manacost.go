package model

import (
	"strconv"
	"strings"
)

// ManaCost is a cost which can be paid to resolve a non-land card.
type ManaCost struct {
	ColorRequirements []ManaColor `json:"colorRequirements"`
	GenericCost       int         `json:"genericCost"`
}

// DeepCopy Creates a deep copy of the ManaCost.
func (m *ManaCost) DeepCopy() ManaCost {
	requirements := make([]ManaColor, len(m.ColorRequirements))
	for i, r := range m.ColorRequirements {
		requirements[i] = r
	}
	return ManaCost{
		ColorRequirements: requirements,
		GenericCost:       m.GenericCost,
	}
}

// ToString Represents the ManaCost as a String.
func (m *ManaCost) ToString() string {
	var stringReqs []string
	for _, requirement := range m.ColorRequirements {
		stringReqs = append(stringReqs, string(requirement))
	}
	return strings.Join(stringReqs, "+") + "+" + strconv.Itoa(m.GenericCost)
}

// GetRemainingCost Computes the total cost which is remaining for this mana cost.
func (m *ManaCost) GetRemainingCost() int {
	return len(m.ColorRequirements) + m.GenericCost
}

// ManaColor Represents a color of mana in the game.
type ManaColor string

const (
	// White Represents the White color of mana.
	White ManaColor = "White"

	// Blue Represents the Blue color of mana.
	Blue ManaColor = "Blue"

	// Black Represents the Black color of mana.
	Black ManaColor = "Black"

	// Red Represents the Red color of mana.
	Red ManaColor = "Red"

	// Green Represents the Green color of mana.
	Green ManaColor = "Green"

	// Colorless Represents the Colorless mana.
	Colorless ManaColor = "Colorless"

	// Whatever Represents any color of mana. This is used primarily for wildcard mana producers, but I'm unsure if this is really necessary.
	Whatever ManaColor = "Whatever"
)
