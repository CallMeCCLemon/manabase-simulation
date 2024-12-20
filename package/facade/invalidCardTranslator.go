package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToExternalInvalidCard converts an InvalidCard from the internal model to the API definition
func ToExternalInvalidCard(invalidCard model.InvalidCard) *api.InvalidCard {
	return &api.InvalidCard{
		Name:   invalidCard.Name,
		Reason: invalidCard.Reason,
	}
}

// ToInternalInvalidCard converts an InvalidCard from the API definition to the internal model
func ToInternalInvalidCard(invalidCard *api.InvalidCard) model.InvalidCard {
	return model.InvalidCard{
		Name:   invalidCard.Name,
		Reason: invalidCard.Reason,
	}
}
