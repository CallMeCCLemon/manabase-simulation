package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToExternalDeckStats converts a DeckStats from the internal model to the API definition.
func ToExternalDeckStats(deckStats model.DeckStats) *api.DeckStats {
	return &api.DeckStats{
		TotalCards:    deckStats.TotalCards,
		Lands:         deckStats.Lands,
		NonLands:      deckStats.NonLands,
		TotalManaPips: ToExternalManaCost(&deckStats.TotalManaPips),
	}
}
