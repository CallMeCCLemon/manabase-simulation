package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToExternalDeckStats converts a DeckStats from the internal model to the API definition.
func ToExternalDeckStats(deckStats model.DeckStats) *api.DeckStats {
	return &api.DeckStats{
		TotalCards:    int32(deckStats.TotalCards),
		Lands:         int32(deckStats.Lands),
		NonLands:      int32(deckStats.NonLands),
		TotalManaPips: ToExternalSimplifiedManaCostFromSimplified(&deckStats.TotalManaPips),
	}
}
