package simulation

import "manabase-simulation/package/model"

// GetDeckStats returns a DeckStats object for a given deck list.
func GetDeckStats(deckList model.DeckList) model.DeckStats {
	return model.DeckStats{
		TotalCards:    int32(len(deckList.Cards)),
		Lands:         int32(len(deckList.Lands)),
		NonLands:      int32(len(deckList.NonLands)),
		TotalManaPips: model.ManaCost{},
	}
}
