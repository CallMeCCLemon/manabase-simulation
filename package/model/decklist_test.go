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

var _ = Describe("Land", func() {
	When("Paying the untapped cost", func() {
		It("Pays the untapped cost correctly", func() {
			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: ShockLand,
				},
				Quantity: 1,
			}

			boardState := BoardState{
				Life: 3,
			}
			Expect(land.PayUntappedCost(&boardState)).To(BeNil())
			Expect(boardState.Life).To(Equal(1))
		})

		It("Doesn't pay the untapped cost when it's not enough life", func() {
			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: ShockLand,
				},
				Quantity: 1,
			}

			boardState := BoardState{
				Life: 2,
			}
			Expect(land.PayUntappedCost(&boardState)).To(HaveOccurred())
			Expect(boardState.Life).To(Equal(2))
		})

		It("Doesn't pay the untapped cost when it's not tapped", func() {
			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: false,
				UntappedCondition: &UntappedCondition{
					Type: ShockLand,
				},
				Quantity: 1,
			}

			boardState := BoardState{
				Life: 2,
			}
			Expect(land.PayUntappedCost(&boardState)).ToNot(HaveOccurred())
			Expect(boardState.Life).To(Equal(2))
		})
	})
})
