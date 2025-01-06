package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
	"manabase-simulation/package/reader"
	"manabase-simulation/package/simulation"
)

var _ = Describe("Decklist JSON Parsing Functions", func() {
	When("Reading a JSON Decklist File", func() {
		deck, err := reader.ReadJSONFile[model.DeckList]("../fixtures/sample_deck.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all of the lands and nonlands", func() {
			Expect(deck.Cards).To(HaveLen(10))
		})

		It("Has the right card count", func() {
			Expect(simulation.GetTotalCardCount(&deck)).To(Equal(10))
		})
	})

	When("Reading the lotus field JSON File", func() {
		deck, err := reader.ReadJSONFile[model.DeckList]("../fixtures/lotus-field-deck.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all of the lands and nonlands", func() {
			Expect(deck.Cards).To(HaveLen(10))
		})

		It("Has the right card count", func() {
			Expect(simulation.GetTotalCardCount(&deck)).To(Equal(60))
		})
	})
})

var _ = Describe("ReadGameConfigJSON", func() {
	When("Reading a JSON File", func() {
		gameConfig, err := reader.ReadJSONFile[model.GameConfiguration]("../fixtures/default-game-config.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all configurations", func() {
			Expect(gameConfig.CardsDrawnPerTurn).To(Equal(1))
			Expect(gameConfig.InitialHandSize).To(Equal(7))
		})
	})
})
