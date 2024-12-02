package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToInternalGameConfiguration translates external API definition to the internal model.
func ToInternalGameConfiguration(gameConfig *api.GameConfiguration) model.GameConfiguration {
	return model.GameConfiguration{
		InitialHandSize:   int(gameConfig.InitialHandSize),
		CardsDrawnPerTurn: int(gameConfig.CardsDrawnPerTurn),
		OnThePlay:         gameConfig.OnThePlay,
	}
}
