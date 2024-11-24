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

var _ = Describe("ShockLand", func() {
	When("Determining if the land can enter untapped", func() {
		It("Can enter untapped when there is 2 life", func() {
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
			Expect(land.CanEnterUntapped(boardState)).To(BeTrue())
		})

		It("Can NOT enter untapped when there is not enough life", func() {
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
				Life: 1,
			}
			Expect(land.CanEnterUntapped(boardState)).To(BeFalse())
		})
	})

	When("Paying the untapped cost", func() {
		It("Pays the cost correctly", func() {
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

		It("Doesn't pay the untapped cost when there is not enough life", func() {
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

var _ = Describe("FastLand", func() {
	When("Determining if the land can enter untapped", func() {
		It("Can enter untapped when there are 2 or less lands", func() {
			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: FastLand,
				},
				Quantity: 1,
			}

			boardState := BoardState{
				Life: 3,
				Lands: []Land{
					*CreateUntappedLand([]ManaColor{White, Blue}),
					*CreateUntappedLand([]ManaColor{White, Blue}),
				},
			}
			Expect(land.CanEnterUntapped(boardState)).To(BeTrue())
		})

		It("Can NOT enter untapped when there are 3 or more lands", func() {
			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: FastLand,
				},
				Quantity: 1,
			}

			boardState := BoardState{
				Life: 3,
				Lands: []Land{
					*CreateUntappedLand([]ManaColor{White, Blue}),
					*CreateUntappedLand([]ManaColor{White, Blue}),
					*CreateUntappedLand([]ManaColor{White, Blue}),
				},
			}
			Expect(land.CanEnterUntapped(boardState)).To(BeFalse())
		})
	})

	When("Paying the untapped cost", func() {
		It("Pays the cost correctly", func() {
			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: FastLand,
				},
				Quantity: 1,
			}

			boardState := BoardState{
				Life: 3,
			}
			Expect(land.PayUntappedCost(&boardState)).ToNot(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})

		It("Doesn't pay the untapped cost when there are too many lands", func() {
			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: FastLand,
				},
				Quantity: 1,
			}

			boardState := BoardState{
				Life: 3,
				Lands: []Land{
					*CreateUntappedLand([]ManaColor{White, Blue}),
					*CreateUntappedLand([]ManaColor{White, Blue}),
					*CreateUntappedLand([]ManaColor{White, Blue}),
				},
			}
			Expect(land.PayUntappedCost(&boardState)).To(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})
	})
})

var _ = Describe("CheckLand", func() {
	When("Determining if the land can enter untapped", func() {
		It("Can enter untapped when there is a land of the right type", func() {
			c := CheckLandData{Plains, Forest}
			cString, _ := c.ToString()

			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: CheckLand,
					Data: &cString,
				},
				Quantity: 1,
			}

			l := CreateUntappedLand([]ManaColor{White, Blue})
			l.Types = []LandType{Plains}

			boardState := BoardState{
				Life: 3,
				Lands: []Land{
					*l,
					*CreateUntappedLand([]ManaColor{White, Blue}),
				},
			}
			Expect(land.CanEnterUntapped(boardState)).To(BeTrue())
		})

		It("Can NOT enter untapped when there is no land of the right type", func() {
			c := CheckLandData{Plains, Forest}
			cString, _ := c.ToString()

			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: CheckLand,
					Data: &cString,
				},
				Quantity: 1,
			}

			l := CreateUntappedLand([]ManaColor{White, Blue})
			l.Types = []LandType{Swamp}

			boardState := BoardState{
				Life: 3,
				Lands: []Land{
					*l,
					*CreateUntappedLand([]ManaColor{White, Blue}),
				},
			}
			Expect(land.CanEnterUntapped(boardState)).To(BeFalse())
		})
	})

	When("Paying the untapped cost", func() {
		It("Pays the cost correctly", func() {
			c := CheckLandData{Plains, Forest}
			cString, _ := c.ToString()

			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: CheckLand,
					Data: &cString,
				},
				Quantity: 1,
			}

			l := CreateUntappedLand([]ManaColor{White, Blue})
			l.Types = []LandType{Plains}

			boardState := BoardState{
				Life: 3,
				Lands: []Land{
					*l,
					*CreateUntappedLand([]ManaColor{White, Blue}),
				},
			}
			Expect(land.PayUntappedCost(&boardState)).ToNot(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})

		It("Doesn't pay the untapped cost when there is no land of the right type", func() {
			c := CheckLandData{Plains, Forest}
			cString, _ := c.ToString()

			land := Land{
				Name:         "test-land-1",
				Colors:       []ManaColor{White, Blue},
				EntersTapped: true,
				UntappedCondition: &UntappedCondition{
					Type: CheckLand,
					Data: &cString,
				},
				Quantity: 1,
			}

			l := CreateUntappedLand([]ManaColor{White, Blue})
			l.Types = []LandType{Swamp}

			boardState := BoardState{
				Life: 3,
				Lands: []Land{
					*l,
					*CreateUntappedLand([]ManaColor{White, Blue}),
				},
			}
			Expect(land.PayUntappedCost(&boardState)).To(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})
	})
})
