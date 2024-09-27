package main

import . "github.com/onsi/ginkgo/v2"

var _ = Describe("Lotus-field", func() {
	deckList, _ := ReadDeckListJSON("./fixtures/lotus-field-deck.json")
	gameConfig, _ := ReadGameConfigJSON("./fixtures/default-game-config.json")

	When("Running the simulation once", func() {

		It("Should determine correctly whether or not the hand generated the right mana", func() {
			SimulateDeck(deckList, gameConfig)
		})
	})
})
