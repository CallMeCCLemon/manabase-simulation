package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
	"manabase-simulation/package/reader"
)

var _ = Describe("Decklist JSON Parsing Functions", func() {
	When("Reading a JSON Decklist File", func() {
		deck, err := reader.ReadJSONFile[model.DeckList]("../fixtures/sample_deck.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all of the lands", func() {
			Expect(deck.Lands).To(HaveLen(5))
		})

		It("Correctly parses all of the non-lands", func() {
			Expect(deck.NonLands).To(HaveLen(5))
		})

		It("Has the right card count", func() {
			Expect(deck.GetTotalCardCount()).To(Equal(10))
		})
	})

	When("Reading the lotus field JSON File", func() {
		deck, err := reader.ReadJSONFile[model.DeckList]("../fixtures/lotus-field-deck.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all of the lands", func() {
			Expect(deck.Lands).To(HaveLen(9))
		})

		It("Correctly parses all of the non-lands", func() {
			Expect(deck.NonLands).To(HaveLen(1))
		})

		It("Has the right card count", func() {
			Expect(deck.GetTotalCardCount()).To(Equal(60))
		})
	})
})

var _ = Describe("ReadGameConfigJSON", func() {
	When("Reading a JSON File", func() {
		gameConfig, err := reader.ReadJSONFile[GameConfiguration]("../fixtures/default-game-config.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all configurations", func() {
			Expect(gameConfig.CardsDrawnPerTurn).To(Equal(1))
			Expect(gameConfig.InitialHandSize).To(Equal(7))
		})
	})
})

var _ = Describe("DeckSimulation", func() {
	var deck model.DeckList
	var gameConfig GameConfiguration
	var objective model.TestObjective

	BeforeEach(func() {
		deck, _ = reader.ReadJSONFile[model.DeckList]("../fixtures/sample_deck.json")
		gameConfig, _ = reader.ReadJSONFile[GameConfiguration]("../fixtures/default-game-config.json")
		objective = model.TestObjective{
			TargetTurn: 3,
			ManaCosts: []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{},
					GenericCost:       1,
				},
			},
		}
	})

	When("Simulating a deck", func() {
		It("Logs correctly to stdout", func() {
			result := SimulateDeck(deck, gameConfig, objective)
			Expect(result).To(BeTrue())
		})
	})

	When("Simulating Lotus Field", func() {
		deck, err := reader.ReadJSONFile[model.DeckList]("../fixtures/lotus-field-deck.json")
		Expect(err).ToNot(HaveOccurred())
		gameConfig, _ := reader.ReadJSONFile[GameConfiguration]("../fixtures/default-game-config.json")
		objective := model.TestObjective{
			TargetTurn: 3,
			ManaCosts: []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White, model.White},
					GenericCost:       1,
				},
			},
		}
		It("Eventually Lotus field succeeds", func() {
			Eventually(func() bool {
				return SimulateDeck(deck, gameConfig, objective)
			}, 3).Should(BeTrue())
		})
	})
})
