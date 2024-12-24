package model

import (
	"github.com/google/go-cmp/cmp"
	"sort"
)

func CreateUntappedLandCard(colors []ManaColor) *Card {
	return &Card{
		Name:     "dummy-untapped-land",
		Land:     CreateUntappedLand(colors),
		NonLand:  nil,
		Quantity: 1,
	}
}

func CreateUntappedLand(colors []ManaColor) *Land {
	return &Land{
		Name:           "dummy-untapped-land",
		Colors:         colors,
		EntersTapped:   false,
		ActivationCost: nil,
		Quantity:       1,
	}
}

func CreateTappedLandCard(colors []ManaColor) *Card {
	return &Card{
		Name:     "dummy-tapped-land",
		Land:     CreateTappedLand(colors),
		NonLand:  nil,
		Quantity: 1,
	}
}

func CreateTappedLand(colors []ManaColor) *Land {
	return &Land{
		Name:           "dummy-tapped-land",
		Colors:         colors,
		EntersTapped:   true,
		ActivationCost: nil,
		Quantity:       1,
	}
}

func CreateSampleNonLandCard() *Card {
	return &Card{
		Name:     "dummy-nonland",
		Land:     nil,
		NonLand:  CreateSampleNonLand(),
		Quantity: 1,
	}
}

func CreateSampleNonLand() *NonLand {
	return &NonLand{
		Name: "dummy-nonland",
		CastingCost: ManaCost{
			ColorRequirements: []ManaColor{White},
			GenericCost:       1,
		},
		Quantity: 1,
	}
}

func CreateManaCost(colors []ManaColor, genericCost int) ManaCost {
	return ManaCost{
		ColorRequirements: colors,
		GenericCost:       genericCost,
	}
}

// CompareDecks compares two different decks to one another.
func CompareDecks(a Deck, b Deck) bool {
	if len(a.Cards) != len(b.Cards) {
		return false
	}

	for i, _ := range a.Cards {
		if !cmp.Equal(a.Cards[i], b.Cards[i]) {
			return false
		}
	}
	return true
}

// SortLandsByRestrictiveness Sorts a list of lands by the number of colors they produce.
func SortLandsByRestrictiveness(lands []Land) []Land {
	sort.Slice(lands, func(i, j int) bool {
		return len(lands[i].Colors) < len(lands[j].Colors)
	})

	return lands
}

// IndexOf finds the index of a specific value in a slice. If not found, returns -1.
func IndexOf[T comparable](slice []T, value T) int {
	for index, v := range slice {
		if v == value {
			return index
		}
	}
	return -1
}
