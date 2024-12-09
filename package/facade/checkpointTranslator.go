package facade

import (
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

// ToExternalResultCheckpoint converts a ResultCheckpoint from the internal model to the API definition
func ToExternalResultCheckpoint(checkpoint model.ResultCheckpoint) *api.ResultCheckpoint {
	return &api.ResultCheckpoint{
		Iterations: checkpoint.Iterations,
		Successes:  checkpoint.Successes,
	}
}
