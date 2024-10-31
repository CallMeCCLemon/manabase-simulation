package main

import (
	"encoding/json"
	"slices"
)

type DeckList struct {
	Lands    []Land    `json:"lands"`
	NonLands []NonLand `json:"nonLands"`
}

func NewDeckList() *DeckList {
	return &DeckList{
		Lands:    []Land{},
		NonLands: []NonLand{},
	}
}

func (d *DeckList) toString() string {
	jsonPayload, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(jsonPayload)
}

func (d *DeckList) getTotalCardCount() int {
	count := 0
	for _, l := range d.Lands {
		count += l.Quantity
	}

	for _, n := range d.NonLands {
		count += n.Quantity
	}

	return count
}

type Land struct {
	Name           string      `json:"name"`
	Colors         []ManaColor `json:"colors"`
	EntersTapped   bool        `json:"entersTapped"`
	ActivationCost []string    `json:"activationCost"`
	Quantity       int         `json:"quantity"`
}

func (l *Land) Equals(land Land) bool {
	// TODO: Include other fields.
	if l.EntersTapped != land.EntersTapped {
		return false
	}
	if l.Name != land.Name {
		return false
	}
	if len(l.Colors) != len(land.Colors) {
		return false
	}
	for _, color := range l.Colors {
		if !slices.Contains(land.Colors, color) {
			return false
		}
	}

	return true
}

type NonLand struct {
	Name        string   `json:"name"`
	CastingCost []string `json:"castingCost"`
	Quantity    int      `json:"quantity"`
}
