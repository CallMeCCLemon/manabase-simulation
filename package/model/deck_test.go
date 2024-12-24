package model

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/reader"
)

var _ = Describe("Deck", func() {
	var deckList DeckList
	var err error

	BeforeEach(func() {
		deckList, err = reader.ReadJSONFile[DeckList]("../../fixtures/sample_deck.json")
		Expect(err).ToNot(HaveOccurred())
	})

	When("Instantiating it from a decklist", func() {
		It("Has 10 cards", func() {
			deck := deckList.GenerateDeck()
			Expect(deck.Cards).To(HaveLen(10))
		})
	})

	When("DeepCopying a deck", func() {
		It("Correctly deep copies", func() {
			deck := deckList.GenerateDeck()
			copiedDeck := deck.DeepCopy()
			Expect(CompareDecks(deck, copiedDeck)).To(BeTrue())
		})
	})

	When("Shuffling the deck", func() {
		It("Randomizes the deck", func() {
			deck := deckList.GenerateDeck()
			unshuffledDeck := deck.DeepCopy()
			Expect(CompareDecks(deck, unshuffledDeck)).To(BeTrue())
			deck.Shuffle()
			Expect(CompareDecks(deck, unshuffledDeck)).To(BeFalse())
		})
	})

	When("Drawing a card from the deck", func() {
		var deck Deck
		var hand Deck
		var firstCard Card

		BeforeEach(func() {
			deck = deckList.GenerateDeck()
			hand = NewDeck()
			firstCard = deck.Cards[0]
			Expect(hand.Cards).To(HaveLen(0))
			Expect(deck.Cards).To(HaveLen(10))
		})

		It("Adds the first card to the hand and Removes the first card from the deck", func() {
			hand = deck.DrawCard(hand)
			Expect(hand.Cards).To(HaveLen(1))
			Expect(hand.Cards[0]).To(Equal(firstCard))
			Expect(deck.Cards).To(HaveLen(9))
			Expect(deck.Cards[0]).ToNot(Equal(firstCard))
		})
	})
})
