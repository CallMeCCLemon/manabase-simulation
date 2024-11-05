package main

import "github.com/google/go-cmp/cmp"

func createUntappedLand(colors []ManaColor) *Land {
	return &Land{
		Name:           "dummy-untapped-land",
		Colors:         colors,
		EntersTapped:   false,
		ActivationCost: nil,
		Quantity:       1,
	}
}

func createTappedLand(colors []ManaColor) *Land {
	return &Land{
		Name:           "dummy-tapped-land",
		Colors:         colors,
		EntersTapped:   true,
		ActivationCost: nil,
		Quantity:       1,
	}
}

func createSampleNonLand() *NonLand {
	return &NonLand{
		Name:        "dummy-nonland",
		CastingCost: []string{"white"},
		Quantity:    1,
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
