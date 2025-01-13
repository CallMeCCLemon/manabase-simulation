package simulation

import "manabase-simulation/package/model"

// GetDeckStats returns a DeckStats object for a given deck list.
func GetDeckStats(deckList model.DeckList) model.DeckStats {
	landCount := 0
	nonLandCount := 0
	cardCount := 0

	spellManaCosts := model.SimplifiedManaCost{}
	landManaProduction := model.SimplifiedManaCost{}

	for _, c := range deckList.Cards {
		cardCount += c.Quantity
		if c.Land != nil {
			landCount += c.Quantity
			for _, color := range c.Land.Colors {
				switch color {
				case model.White:
					landManaProduction.WhiteMana += c.Quantity
				case model.Blue:
					landManaProduction.BlueMana += c.Quantity
				case model.Black:
					landManaProduction.BlackMana += c.Quantity
				case model.Red:
					landManaProduction.RedMana += c.Quantity
				case model.Green:
					landManaProduction.GreenMana += c.Quantity
				case model.Colorless:
					landManaProduction.ColorlessMana += c.Quantity
				}
			}
		} else {
			nonLandCount += c.Quantity
			spellManaCosts.GenericMana += c.Quantity * c.NonLand.CastingCost.GenericCost
			for _, color := range c.NonLand.CastingCost.ColorRequirements {
				switch color {
				case model.White:
					spellManaCosts.WhiteMana += c.Quantity
				case model.Blue:
					spellManaCosts.BlueMana += c.Quantity
				case model.Black:
					spellManaCosts.BlackMana += c.Quantity
				case model.Red:
					spellManaCosts.RedMana += c.Quantity
				case model.Green:
					spellManaCosts.GreenMana += c.Quantity
				case model.Colorless:
					spellManaCosts.ColorlessMana += c.Quantity
				case model.Azorius:
					spellManaCosts.WhiteMana += c.Quantity
					spellManaCosts.BlueMana += c.Quantity
				case model.Orzhov:
					spellManaCosts.WhiteMana += c.Quantity
					spellManaCosts.BlackMana += c.Quantity
				case model.Boros:
					spellManaCosts.WhiteMana += c.Quantity
					spellManaCosts.RedMana += c.Quantity
				case model.Selesnya:
					spellManaCosts.WhiteMana += c.Quantity
					spellManaCosts.GreenMana += c.Quantity
				case model.Dimir:
					spellManaCosts.BlueMana += c.Quantity
					spellManaCosts.BlackMana += c.Quantity
				case model.Izzet:
					spellManaCosts.BlueMana += c.Quantity
					spellManaCosts.RedMana += c.Quantity
				case model.Simic:
					spellManaCosts.BlueMana += c.Quantity
					spellManaCosts.GreenMana += c.Quantity
				case model.Rakdos:
					spellManaCosts.BlackMana += c.Quantity
					spellManaCosts.RedMana += c.Quantity
				case model.Golgari:
					spellManaCosts.RedMana += c.Quantity
					spellManaCosts.GreenMana += c.Quantity
				case model.Gruul:
					spellManaCosts.GreenMana += c.Quantity
					spellManaCosts.RedMana += c.Quantity
				}
			}
		}
	}

	return model.DeckStats{
		TotalCards:    cardCount,
		Lands:         landCount,
		NonLands:      nonLandCount,
		TotalManaPips: spellManaCosts,
		LandStats: model.LandStats{
			LandManaProduction: landManaProduction,
		},
	}
}
