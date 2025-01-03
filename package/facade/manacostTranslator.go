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
