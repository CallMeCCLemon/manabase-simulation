package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToInternalDeckList converts a DeckList from the API definition to the internal model
func ToInternalDeckList(decklist *api.DeckList) model.DeckList {
	lands := make([]model.Land, len(decklist.Lands))
	for i, l := range decklist.Lands {
		lands[i] = toInternalLand(l)
	}

	nonLands := make([]model.NonLand, len(decklist.NonLands))
	for i, n := range decklist.NonLands {
		nonLands[i] = toInternalNonLand(n)
	}

	return model.DeckList{
		Lands:    lands,
		NonLands: nonLands,
	}
}

func toInternalLand(land *api.Land) model.Land {
	types := make([]model.LandType, len(land.Types))
	for i, t := range land.Types {
		types[i] = toInternalLandType(t)
	}

	colors := make([]model.ManaColor, len(land.Colors))
	for i, c := range land.Colors {
		colors[i] = toInternalManaColor(c)
	}

	return model.Land{
		Name:              land.Name,
		Types:             types,
		Colors:            colors,
		EntersTapped:      land.EntersTapped,
		ActivationCost:    toInternalActivationCost(land.ActivationCost),
		UntappedCondition: toInternalUntappedCondition(land.UntappedCondition),
		Quantity:          int(land.Quantity),
	}
}

func toInternalLandType(landType api.LandType) model.LandType {
	switch landType {
	case api.LandType_FOREST:
		return model.Forest
	case api.LandType_ISLAND:
		return model.Island
	case api.LandType_MOUNTAIN:
		return model.Mountain
	case api.LandType_SWAMP:
		return model.Swamp
	case api.LandType_PLAINS:
		return model.Plains
	default:
		// API Definition should prevent this.
		panic("Unknown land type")
	}
}

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

func toInternalNonLand(nonLand *api.NonLand) model.NonLand {
	return model.NonLand{
		Name:        nonLand.Name,
		CastingCost: *toInternalManaCost(nonLand.CastingCost),
		Quantity:    int(nonLand.Quantity),
	}
}

func toInternalActivationCost(activationCost *api.ActivationCost) *model.ActivationCost {
	if activationCost == nil {
		return nil
	}

	life := int(activationCost.Life)
	return &model.ActivationCost{
		Life:     &life,
		ManaCost: toInternalManaCost(activationCost.ManaCost),
	}
}

func toInternalUntappedCondition(untappedCondition *api.UntappedCondition) *model.UntappedCondition {
	if untappedCondition == nil {
		return nil
	}

	return &model.UntappedCondition{
		Type: toInternalConditionType(untappedCondition.Type),
		Data: &untappedCondition.Data,
	}
}

func toInternalConditionType(conditionType api.ConditionType) model.ConditionType {
	switch conditionType {
	case api.ConditionType_SHOCK_LAND:
		return model.ShockLand
	case api.ConditionType_FAST_LAND:
		return model.FastLand
	case api.ConditionType_CHECK_LAND:
		return model.CheckLand
	default:
		panic("Unknown condition type")
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
