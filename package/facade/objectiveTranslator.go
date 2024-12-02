package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToInternalTestObjective translates an external API definition to the internal model.
func ToInternalTestObjective(objective *api.Objective) model.TestObjective {
	manaCosts := make([]model.ManaCost, len(objective.ManaCosts))
	for i, m := range objective.ManaCosts {
		manaCosts[i] = *toInternalManaCost(m)
	}

	return model.TestObjective{
		TargetTurn: int(objective.TargetTurn),
		ManaCosts:  manaCosts,
	}
}
