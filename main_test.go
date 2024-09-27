package main

import (
	"github.com/google/go-cmp/cmp"
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
			Expect(deck.getTotalCardCount()).To(Equal(2))
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

	BeforeEach(func() {
		deck, _ = ReadDeckListJSON("./fixtures/sample_deck.json")
		gameConfig, _ = ReadGameConfigJSON("./fixtures/default-game-config.json")
	})

	When("Simulating a deck", func() {
		It("Logs correctly to stdout", func() {
			SimulateDeck(deck, gameConfig)
		})
	})
})

// CompareDecks compares two different decks to one another.
func CompareDecks(a Deck, b Deck) bool {
	if len(a.Cards) != len(b.Cards) {
		return false
	}

	for i, _ := range a.Cards {
		if !cmp.Equal(a.Cards[i], b.Cards[i]) {
			return false
		}
	}
	return true
}
