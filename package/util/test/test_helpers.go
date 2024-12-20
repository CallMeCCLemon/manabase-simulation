// Package test contains helper functions for testing.
package test

import (
	"fmt"
	"gorm.io/driver/postgres"
	"manabase-simulation/package/model"
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// NewLandCard creates a new land card instance for testing.
func NewLandCard(name string) *model.Card {
	return &model.Card{
		Name:    fmt.Sprintf("%s-%s", name, randSeq(8)),
		Land:    NewLand(name),
		NonLand: nil,
	}
}

func NewLand(name string) *model.Land {
	life := 3
	return &model.Land{
		Name: fmt.Sprintf("%s-%s", name, randSeq(8)),
		Colors: []model.ManaColor{
			model.White,
			model.Blue,
		},
		EntersTapped: true,
		UntappedCondition: &model.UntappedCondition{
			Type: model.ShockLand,
		},
		Quantity: 1,
		ActivationCost: &model.ActivationCost{
			Life: &life,
			ManaCost: &model.ManaCost{
				ColorRequirements: []model.ManaColor{
					model.White,
					model.White,
				},
				GenericCost: 1,
			},
		},
	}
}

func NewNonLandCard(name string) *model.Card {
	return &model.Card{
		Name:    fmt.Sprintf("%s-%s", name, randSeq(8)),
		Land:    nil,
		NonLand: NewNonLand(name),
	}
}

func NewNonLand(name string) *model.NonLand {
	return &model.NonLand{
		Name: fmt.Sprintf("%s-%s", name, randSeq(8)),
		CastingCost: model.ManaCost{
			ColorRequirements: []model.ManaColor{
				model.White,
				model.White,
				model.White,
			},
			GenericCost: 1,
		},
		Quantity: 1,
	}
}

func GetDBConfig() postgres.Config {
	cfg := postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "app", os.Getenv("PORT")),
	}
	return cfg
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
