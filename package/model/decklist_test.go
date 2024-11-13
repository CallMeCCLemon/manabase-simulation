package model

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckList", func() {
	When("Getting the card count", func() {
		It("correctly counts all of the lands and nonlands", func() {
			deckList := DeckList{
				Lands: []Land{
					*CreateUntappedLand([]ManaColor{White}),
					*CreateUntappedLand([]ManaColor{White, Blue}),
					*CreateUntappedLand([]ManaColor{Red, Blue}),
					*CreateUntappedLand([]ManaColor{Red, Blue, White}),
				},
				NonLands: []NonLand{
					*CreateSampleNonLand(),
					*CreateSampleNonLand(),
					*CreateSampleNonLand(),
				},
			}

			Expect(deckList.GetTotalCardCount()).To(Equal(7))
		})
	})
})
