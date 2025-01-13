package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToExternalDeckStats converts a DeckStats from the internal model to the API definition.
func ToExternalDeckStats(deckStats model.DeckStats) *api.DeckStats {
	landStats := ToExternalLandStats(&deckStats.LandStats)

	return &api.DeckStats{
		TotalCards:    int32(deckStats.TotalCards),
		Lands:         int32(deckStats.Lands),
		NonLands:      int32(deckStats.NonLands),
		TotalManaPips: ToExternalSimplifiedManaCostFromSimplified(&deckStats.TotalManaPips),
		LandStats:     &landStats,
	}
}

// ToExternalLandStats converts a LandStats object from the internal model to the API definition
func ToExternalLandStats(landStats *model.LandStats) api.LandStats {
	return api.LandStats{
		LandManaProduction: ToExternalSimplifiedManaCostFromSimplified(&landStats.LandManaProduction),
	}
}
