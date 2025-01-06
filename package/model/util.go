package model

import (
	"github.com/google/go-cmp/cmp"
)

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

// IndexOf finds the index of a specific value in a slice. If not found, returns -1.
func IndexOf[T comparable](slice []T, value T) int {
	for index, v := range slice {
		if v == value {
			return index
		}
	}
	return -1
}
