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

func (m *ManaCost) toString() string {
	var stringReqs []string
	for _, requirement := range m.ColorRequirements {
		stringReqs = append(stringReqs, string(requirement))
	}
	return strings.Join(stringReqs, "+") + "+" + strconv.Itoa(m.GenericCost)
}
