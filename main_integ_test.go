package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lotus-field", func() {
	deckList, _ := ReadDeckListJSON("./fixtures/lotus-field-deck.json")
	gameConfig, _ := ReadGameConfigJSON("./fixtures/default-game-config.json")
	objective := TestObjective{
		TargetTurn: 3,
		ManaCosts: []ManaCost{
			{
				ColorRequirements: []ManaColor{white, white},
				GenericCost:       1,
			},
		},
	}

	When("Running the simulation once", func() {
		It("Should determine correctly whether or not the hand generated the right mana", func() {
			success := SimulateDeck(deckList, gameConfig, objective)
			Expect(success).To(BeTrue())
		})
	})
})
