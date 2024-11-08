package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deck", func() {
	deckList, _ := ReadJSONFile[DeckList]("./fixtures/sample_deck.json")

	When("Instantiating it from a decklist", func() {
		deck := deckList.GenerateDeck()

		It("Has 10 cards", func() {
			Expect(deck.Cards).To(HaveLen(10))
		})
	})

	When("DeepCopying a deck", func() {
		deck := deckList.GenerateDeck()
		copiedDeck := deck.DeepCopy()

		It("Correctly deep copies", func() {
			Expect(CompareDecks(deck, copiedDeck)).To(BeTrue())
		})
	})

	When("Shuffling the deck", func() {
		deck := deckList.GenerateDeck()
		unshuffledDeck := deck.DeepCopy()
		Expect(CompareDecks(deck, unshuffledDeck)).To(BeTrue())

		It("Randomizes the deck", func() {
			deck.Shuffle()
			Expect(CompareDecks(deck, unshuffledDeck)).To(BeFalse())
		})
	})

	When("Drawing a card from the deck", func() {
		deck := deckList.GenerateDeck()
		hand := NewDeck()
		firstCard := deck.Cards[0]
		It("Has 10 cards", func() {
			Expect(hand.Cards).To(HaveLen(0))
			Expect(deck.Cards).To(HaveLen(10))
		})

		It("Adds the first card to the hand", func() {
			hand = deck.DrawCard(hand)
			Expect(hand.Cards).To(HaveLen(1))
			Expect(hand.Cards[0]).To(Equal(firstCard))
		})

		It("Removes the first card from the deck", func() {
			Expect(deck.Cards).To(HaveLen(9))
			Expect(deck.Cards[0]).ToNot(Equal(firstCard))
		})
	})
})
