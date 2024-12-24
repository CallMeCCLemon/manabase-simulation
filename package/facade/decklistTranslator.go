package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

func toInternalManaColor(manaColor api.ManaColor) model.ManaColor {
	switch manaColor {
	case api.ManaColor_WHITE:
		return model.White
	case api.ManaColor_BLUE:
		return model.Blue
	case api.ManaColor_BLACK:
		return model.Black
	case api.ManaColor_RED:
		return model.Red
	case api.ManaColor_GREEN:
		return model.Green
	case api.ManaColor_COLORLESS:
		return model.Colorless
	default:
		// API Definition should prevent this.
		panic("Unknown mana color")
	}
}

func toInternalManaCost(manaCost *api.ManaCost) *model.ManaCost {
	if manaCost == nil {
		return nil
	}

	colorRequirements := make([]model.ManaColor, len(manaCost.ColorRequirements))
	for i, c := range manaCost.ColorRequirements {
		colorRequirements[i] = toInternalManaColor(c)
	}

	return &model.ManaCost{
		ColorRequirements: colorRequirements,
		GenericCost:       int(manaCost.GenericCost),
	}
}
