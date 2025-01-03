package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

func ToExternalManaCost(manaCost *model.ManaCost) *api.ManaCost {
	if manaCost == nil {
		return nil
	}

	colorRequirements := make([]api.ManaColor, len(manaCost.ColorRequirements))
	for i, c := range manaCost.ColorRequirements {
		colorRequirements[i] = toExternalManaColor(c)
	}

	return &api.ManaCost{
		ColorRequirements: colorRequirements,
		GenericCost:       int32(manaCost.GenericCost),
	}
}

func ToExternalSimplifiedManaCostFromSimplified(cost *model.SimplifiedManaCost) *api.SimplifiedManaCost {
	if cost == nil {
		return nil
	}

	manaCost := api.SimplifiedManaCost{
		GenericCost: int32(cost.GenericMana),
	}

	manaCost.WhiteMana = int32(cost.WhiteMana + cost.AzoriusMana + cost.OrzhovMana + cost.BorosMana + cost.SelesnyaMana)
	manaCost.BlueMana = int32(cost.BlueMana + cost.DimirMana + cost.IzzetMana + cost.SimicMana + cost.AzoriusMana)
	manaCost.BlackMana = int32(cost.BlackMana + cost.RakdosMana + cost.GolgariMana + cost.OrzhovMana + cost.DimirMana)
	manaCost.RedMana = int32(cost.RedMana + cost.RakdosMana + cost.GruulMana + cost.BorosMana + cost.IzzetMana)
	manaCost.GreenMana = int32(cost.GreenMana + cost.GruulMana + cost.SelesnyaMana + cost.SimicMana + cost.GolgariMana)
	manaCost.ColorlessMana = int32(cost.ColorlessMana)

	return &manaCost
}

// ToExternalSimplifiedManaCost converts a ManaCost to an external SimplifiedManaCost.
func ToExternalSimplifiedManaCost(manaCost *model.ManaCost) *api.SimplifiedManaCost {
	if manaCost == nil {
		return nil
	}

	simplifiedManaCost := api.SimplifiedManaCost{
		GenericCost: int32(manaCost.GenericCost),
	}

	for _, c := range manaCost.ColorRequirements {
		switch c {
		case model.White:
			simplifiedManaCost.WhiteMana += int32(1)
		case model.Blue:
			simplifiedManaCost.BlueMana += int32(1)
		case model.Black:
			simplifiedManaCost.BlackMana += int32(1)
		case model.Red:
			simplifiedManaCost.RedMana += int32(1)
		case model.Green:
			simplifiedManaCost.GreenMana += int32(1)
		case model.Colorless:
			simplifiedManaCost.ColorlessMana += int32(1)
		}
	}

	return &simplifiedManaCost
}

func toExternalManaColor(manaColor model.ManaColor) api.ManaColor {
	switch manaColor {
	case model.White:
		return api.ManaColor_WHITE
	case model.Blue:
		return api.ManaColor_BLUE
	case model.Black:
		return api.ManaColor_BLACK
	case model.Red:
		return api.ManaColor_RED
	case model.Green:
		return api.ManaColor_GREEN
	case model.Colorless:
		return api.ManaColor_COLORLESS
	default:
		// API Definition should prevent this.
		panic("Unknown mana color")
	}
}
