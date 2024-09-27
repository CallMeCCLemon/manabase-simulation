package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckList", func() {
	When("Reading a JSON File", func() {
		deck, err := ReadDeckJSON("./fixtures/sample_deck.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all of the lands", func() {
			Expect(deck.Lands).To(HaveLen(1))
		})

		It("Correctly parses all of the non-lands", func() {
			Expect(deck.NonLands).To(HaveLen(1))
		})

		It("Has the right card count", func() {
			Expect(deck.getTotalCardCount()).To(Equal(2))
		})
	})

	When("Reading the lotus field JSON File", func() {
		deck, err := ReadDeckJSON("./fixtures/lotus-field-deck.json")

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
			Expect(deck.getTotalCardCount()).To(Equal(60))
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

	BeforeEach(func() {
		deck, _ = ReadDeckJSON("./fixtures/sample_deck.json")
		gameConfig, _ = ReadGameConfigJSON("./fixtures/default-game-config.json")
	})

	When("Simulating a deck", func() {
		It("Logs correctly to stdout", func() {
			SimulateDeck(deck, gameConfig)
		})
	})
})
