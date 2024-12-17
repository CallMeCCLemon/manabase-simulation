// Package test contains helper functions for testing.
package test

import (
	"fmt"
	"manabase-simulation/package/model"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// NewLandCard creates a new land card instance for testing.
func NewLandCard(name string) *model.Card {
	return &model.Card{
		Name: fmt.Sprintf("%s-%s", name, randSeq(8)),
		Land: &model.Land{
			Name: name,
			Colors: []model.ManaColor{
				model.White,
				model.Blue,
			},
			EntersTapped: true,
			UntappedCondition: &model.UntappedCondition{
				Type: model.ShockLand,
			},
			Quantity: 1,
		},
		NonLand: nil,
	}
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
