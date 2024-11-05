package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckList", func() {
	When("Getting the card count", func() {
		It("correctly counts all of the lands and nonlands", func() {
			deckList := DeckList{
				Lands: []Land{
					*createUntappedLand([]ManaColor{white}),
					*createUntappedLand([]ManaColor{white, blue}),
					*createUntappedLand([]ManaColor{red, blue}),
					*createUntappedLand([]ManaColor{red, blue, white}),
				},
				NonLands: []NonLand{
					*createSampleNonLand(),
					*createSampleNonLand(),
					*createSampleNonLand(),
				},
			}

			Expect(deckList.GetTotalCardCount()).To(Equal(7))
		})
	})
})
