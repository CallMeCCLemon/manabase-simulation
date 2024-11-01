package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckList", func() {
	When("Reading a JSON File", func() {
		deck, err := ReadDeckListJSON("./fixtures/sample_deck.json")

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
			Expect(deck.getTotalCardCount()).To(Equal(10))
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
			Expect(deck.getTotalCardCount()).To(Equal(60))
		})
	})
})

var _ = Describe("Deck", func() {
	deckList, _ := ReadDeckListJSON("./fixtures/lotus-field-deck.json")

	When("Instantiating it from a decklist", func() {
		deck := GenerateDeck(deckList)

		It("Has 60 cards", func() {
			Expect(deck.Cards).To(HaveLen(60))
		})
	})

	When("DeepCopying a deck", func() {
		deck := GenerateDeck(deckList)
		copiedDeck := deck.DeepCopy()

		It("Correctly deep copies", func() {
			Expect(CompareDecks(deck, copiedDeck)).To(BeTrue())
		})
	})

	When("Shuffling the deck", func() {
		deck := GenerateDeck(deckList)
		unshuffledDeck := deck.DeepCopy()
		Expect(CompareDecks(deck, unshuffledDeck)).To(BeTrue())

		It("Randomizes the deck", func() {
			deck.Shuffle()
			Expect(CompareDecks(deck, unshuffledDeck)).To(BeFalse())
		})
	})

	When("Drawing a card from the deck", func() {
		deck := GenerateDeck(deckList)
		hand := NewDeck()
		firstCard := deck.Cards[0]
		Expect(hand.Cards).To(HaveLen(0))
		Expect(deck.Cards).To(HaveLen(60))
		hand = deck.DrawCard(hand)

		It("Adds the first card to the hand", func() {
			Expect(hand.Cards).To(HaveLen(1))
			Expect(hand.Cards[0]).To(Equal(firstCard))
		})

		It("Removes the first card from the deck", func() {
			Expect(deck.Cards).To(HaveLen(59))
			Expect(deck.Cards[0]).ToNot(Equal(firstCard))
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
