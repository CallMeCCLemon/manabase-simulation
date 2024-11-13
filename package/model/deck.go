package model

import "math/rand/v2"

// Deck Represents a deck of cards where each card has a quantity of 1. This is the primary data model used during the simulation.
type Deck struct {
	Cards []Card
}

// NewDeck Creates a new Deck instance.
func NewDeck() Deck {
	return Deck{
		Cards: []Card{},
	}
}

// Shuffle Shuffles the deck.
func (d *Deck) Shuffle() {
	shuffledCards := make([]Card, len(d.Cards))
	perm := rand.Perm(len(d.Cards))

	for i, v := range perm {
		shuffledCards[v] = d.Cards[i]
	}
	d.Cards = shuffledCards
}

// DeepCopy Copies a Deck to a new obj.
func (d *Deck) DeepCopy() Deck {
	newDeck := NewDeck()
	for _, card := range d.Cards {
		newDeck.Cards = append(newDeck.Cards, card)
	}

	return newDeck
}

// DrawCard Draws a card from the deck and adds it to the hand.
func (d *Deck) DrawCard(hand Deck) (updatedHand Deck) {
	hand.Cards = append(hand.Cards, d.Cards[0])
	d.Cards = d.Cards[1:]

	return hand
}
