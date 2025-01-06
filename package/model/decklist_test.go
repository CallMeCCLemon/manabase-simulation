package model

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckList", func() {
	When("Getting the card count", func() {
		It("correctly counts all of the lands and nonlands", func() {
			deckList := DeckList{
				Cards: []Card{
					*CreateUntappedLandCard([]ManaColor{White}),
					*CreateUntappedLandCard([]ManaColor{White, Blue}),
					*CreateUntappedLandCard([]ManaColor{Red, Blue}),
					*CreateUntappedLandCard([]ManaColor{Red, Blue, White}),
					*CreateSampleNonLandCard(),
					*CreateSampleNonLandCard(),
					*CreateSampleNonLandCard(),
				},
			}

			Expect(deckList.GetTotalCardCount()).To(Equal(7))
		})
	})
})
