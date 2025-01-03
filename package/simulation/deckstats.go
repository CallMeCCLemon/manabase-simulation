package simulation

import "manabase-simulation/package/model"

// GetDeckStats returns a DeckStats object for a given deck list.
func GetDeckStats(deckList model.DeckList) model.DeckStats {
	landCount := 0
	nonLandCount := 0
	cardCount := 0

	manaCost := model.SimplifiedManaCost{}

	for _, c := range deckList.Cards {
		cardCount += c.Quantity
		if c.Land != nil {
			landCount += c.Quantity
		} else {
			nonLandCount += c.Quantity
			manaCost.GenericMana += c.Quantity * c.NonLand.CastingCost.GenericCost
			for _, color := range c.NonLand.CastingCost.ColorRequirements {
				switch color {
				case model.White:
					manaCost.WhiteMana += c.Quantity
				case model.Blue:
					manaCost.BlueMana += c.Quantity
				case model.Black:
					manaCost.BlackMana += c.Quantity
				case model.Red:
					manaCost.RedMana += c.Quantity
				case model.Green:
					manaCost.GreenMana += c.Quantity
				case model.Colorless:
					manaCost.ColorlessMana += c.Quantity
				case model.Azorius:
					manaCost.WhiteMana += c.Quantity
					manaCost.BlueMana += c.Quantity
				case model.Orzhov:
					manaCost.WhiteMana += c.Quantity
					manaCost.BlackMana += c.Quantity
				case model.Boros:
					manaCost.WhiteMana += c.Quantity
					manaCost.RedMana += c.Quantity
				case model.Selesnya:
					manaCost.WhiteMana += c.Quantity
					manaCost.GreenMana += c.Quantity
				case model.Dimir:
					manaCost.BlueMana += c.Quantity
					manaCost.BlackMana += c.Quantity
				case model.Izzet:
					manaCost.BlueMana += c.Quantity
					manaCost.RedMana += c.Quantity
				case model.Simic:
					manaCost.BlueMana += c.Quantity
					manaCost.GreenMana += c.Quantity
				case model.Rakdos:
					manaCost.BlackMana += c.Quantity
					manaCost.RedMana += c.Quantity
				case model.Golgari:
					manaCost.RedMana += c.Quantity
					manaCost.GreenMana += c.Quantity
				case model.Gruul:
					manaCost.GreenMana += c.Quantity
					manaCost.RedMana += c.Quantity
				}
			}
		}
	}

	return model.DeckStats{
		TotalCards:    cardCount,
		Lands:         landCount,
		NonLands:      nonLandCount,
		TotalManaPips: manaCost,
	}
}
