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
			*CreateUntappedLandCard([]model.ManaColor{model.White}),
			*CreateUntappedLandCard([]model.ManaColor{model.White, model.Blue}),
			*CreateSampleNonLandCard(),
			*CreateSampleNonLandCard(),
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
		NonLand: NewNonLand(),
	}
}

func NewNonLand() *model.NonLand {
	return &model.NonLand{
		CastingCost: model.ManaCost{
			ColorRequirements: []model.ManaColor{
				model.White,
				model.White,
				model.White,
			},
			GenericCost: 1,
		},
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

func CreateUntappedLandCard(colors []model.ManaColor) *model.Card {
	return &model.Card{
		Name:     "dummy-untapped-land",
		Land:     CreateUntappedLand(colors),
		NonLand:  nil,
		Quantity: 1,
	}
}

func CreateUntappedLand(colors []model.ManaColor) *model.Land {
	return &model.Land{
		Colors:         colors,
		EntersTapped:   false,
		ActivationCost: nil,
	}
}

func CreateTappedLandCard(colors []model.ManaColor) *model.Card {
	return &model.Card{
		Name:     "dummy-tapped-land",
		Land:     CreateTappedLand(colors),
		NonLand:  nil,
		Quantity: 1,
	}
}

func CreateTappedLand(colors []model.ManaColor) *model.Land {
	return &model.Land{
		Colors:         colors,
		EntersTapped:   true,
		ActivationCost: nil,
	}
}

func CreateSampleNonLandCard() *model.Card {
	return &model.Card{
		Name:     "dummy-nonland",
		Land:     nil,
		NonLand:  CreateSampleNonLand(),
		Quantity: 1,
	}
}

func CreateSampleNonLand() *model.NonLand {
	return &model.NonLand{
		CastingCost: model.ManaCost{
			ColorRequirements: []model.ManaColor{model.White},
			GenericCost:       1,
		},
	}
}

func CreateManaCost(colors []model.ManaColor, genericCost int) model.ManaCost {
	return model.ManaCost{
		ColorRequirements: colors,
		GenericCost:       genericCost,
	}
}
