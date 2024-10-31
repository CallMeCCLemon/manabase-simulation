package main

import "math/rand/v2"

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	return Deck{
		Cards: []Card{},
	}
}

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

func (d *Deck) DrawCard(hand Deck) (updatedHand Deck) {
	hand.Cards = append(hand.Cards, d.Cards[0])
	d.Cards = d.Cards[1:]

	return hand
}
