package main

import (
	"strconv"
	"strings"
)

type ManaCost struct {
	ColorRequirements []ManaColor `json:"colorRequirements"`
	GenericCost       int         `json:"genericCost"`
}

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
