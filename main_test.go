package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decklist JSON Parsing Functions", func() {
	When("Reading a JSON Decklist File", func() {
		deck, err := ReadDeckListJSON("./fixtures/sample_deck.json")

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
		deck, err := ReadDeckListJSON("./fixtures/lotus-field-deck.json")

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
		gameConfig, err := ReadGameConfigJSON("./fixtures/default-game-config.json")

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
	var deck DeckList
	var gameConfig GameConfiguration
	var objective TestObjective

	BeforeEach(func() {
		deck, _ = ReadDeckListJSON("./fixtures/sample_deck.json")
		gameConfig, _ = ReadGameConfigJSON("./fixtures/default-game-config.json")
		objective = TestObjective{
			TargetTurn: 3,
			ManaCosts: []ManaCost{
				{
					ColorRequirements: []ManaColor{},
					GenericCost:       1,
				},
			},
		}
	})

	When("Simulating a deck", func() {
		It("Logs correctly to stdout", func() {
			SimulateDeck(deck, gameConfig, objective)
		})

	})
})

var _ = Describe("Sorting a list of lands", func() {
	When("Sorting a list of lands with different quantities of mana they can tap for", func() {
		It("Returns a list where each subsequent len of land.Colors >= prevLand.Colors", func() {
			var lands []Land
			lands = append(lands, *createUntappedLand([]ManaColor{white, green, red, black}))
			lands = append(lands, *createUntappedLand([]ManaColor{white}))
			lands = append(lands, *createUntappedLand([]ManaColor{white, green, red}))
			lands = append(lands, *createUntappedLand([]ManaColor{white, blue}))
			lands = append(lands, *createUntappedLand([]ManaColor{white, green, red, black, blue}))

			sortedLands := SortLandsByRestrictiveness(lands)

			Expect(sortedLands[0].Colors).To(HaveLen(1))
			Expect(sortedLands[1].Colors).To(HaveLen(2))
			Expect(sortedLands[2].Colors).To(HaveLen(3))
			Expect(sortedLands[3].Colors).To(HaveLen(4))
			Expect(sortedLands[4].Colors).To(HaveLen(5))
		})
	})
})
