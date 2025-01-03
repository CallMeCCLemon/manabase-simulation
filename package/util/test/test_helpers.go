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

func NewDeckList() *model.DeckList {
	return &model.DeckList{
		Cards: []model.Card{
			*model.CreateUntappedLandCard([]model.ManaColor{model.White}),
			*model.CreateUntappedLandCard([]model.ManaColor{model.White, model.Blue}),
			*model.CreateSampleNonLandCard(),
			*model.CreateSampleNonLandCard(),
		},
	}
}

// NewLandCard creates a new land card instance for testing.
func NewLandCard(name string) *model.Card {
	return &model.Card{
		Name:    fmt.Sprintf("%s-%s", name, randSeq(8)),
		Land:    NewLand(),
		NonLand: nil,
	}
}

func NewLand() *model.Land {
	life := 3
	return &model.Land{
		Colors: []model.ManaColor{
			model.White,
			model.Blue,
		},
		EntersTapped: true,
		UntappedCondition: &model.UntappedCondition{
			Type: model.ShockLand,
		},
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
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("host"), os.Getenv("username"), os.Getenv("password"), "app", os.Getenv("port")),
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
