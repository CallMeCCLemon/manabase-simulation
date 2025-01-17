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

// SimplifiedManaCost is a simplified cost which can be paid to resolve a non-land card.
type SimplifiedManaCost struct {
	WhiteMana     int `json:"whiteMana"`
	BlueMana      int `json:"blueMana"`
	BlackMana     int `json:"blackMana"`
	RedMana       int `json:"redMana"`
	GreenMana     int `json:"greenMana"`
	ColorlessMana int `json:"colorlessMana"`
	AzoriusMana   int `json:"azoriusMana"`
	OrzhovMana    int `json:"orzhovMana"`
	BorosMana     int `json:"borosMana"`
	SelesnyaMana  int `json:"selesnyaMana"`
	DimirMana     int `json:"dimirMana"`
	IzzetMana     int `json:"izzetMana"`
	SimicMana     int `json:"simicMana"`
	RakdosMana    int `json:"rakdosMana"`
	GolgariMana   int `json:"golgariMana"`
	GruulMana     int `json:"gruulMana"`
	GenericMana   int `json:"genericMana"`
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

	// Azorius Represents White and Blue mana hybrid.
	Azorius ManaColor = "Azorius"

	// Orzhov Represents White and Black mana hybrid.
	Orzhov ManaColor = "Orzhov"

	// Boros Represents White and Red mana hybrid.
	Boros ManaColor = "Boros"

	// Selesnya Represents White and Green mana hybrid.
	Selesnya ManaColor = "Selesnya"

	// Dimir Represents Blue and Black mana hybrid.
	Dimir ManaColor = "Dimir"

	// Izzet Represents Blue and Red mana hybrid.
	Izzet ManaColor = "Izzet"

	// Simic Represents Blue and Green mana hybrid.
	Simic ManaColor = "Simic"

	// Rakdos Represents Black and Red mana hybrid.
	Rakdos ManaColor = "Rakdos"

	// Golgari Represents Black and Green mana hybrid.
	Golgari ManaColor = "Golgari"

	// Gruul Represents Red and Green mana hybrid.
	Gruul ManaColor = "Gruul"
)
