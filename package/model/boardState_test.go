package model

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BoardState", func() {
	var boardState BoardState
	BeforeEach(func() {
		boardState = NewBoardState()
	})

	When("Creating a board state", func() {
		It("Has no lands", func() {
			boardState = NewBoardState()
			Expect(boardState.Lands).To(BeEmpty())
		})
	})

	When("Validating a test objective", func() {
		It("Is able to solve for a simple objective with basic lands as met", func() {
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							White,
							White,
							Blue,
						},
						GenericCost: 0,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Is able to solve for an objective with basic and dual lands as met", func() {
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							White,
							White,
							Blue,
						},
						GenericCost: 0,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())

			obj = TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							White,
							Red,
							Blue,
						},
						GenericCost: 0,
					},
				},
			}

			isMet, _ = boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Is able to solve for an objective with basic, dual, and triome lands as met", func() {
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue, Green}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							White,
							White,
							Blue,
							Red,
						},
						GenericCost: 0,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Is able to determine an infeasible solution with 'dead' lands", func() {
			// Dead land is a plains here.
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue, Green}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							Blue,
							Blue,
							Blue,
							Green,
						},
						GenericCost: 0,
					},
				},
			}
			isMet, combos := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeFalse())
			Expect(combos).To(HaveLen(2))
		})

		It("Verify Generic costs are able to be paid by unusable lands", func() {
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue, Green}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							Blue,
							Blue,
							Green,
						},
						GenericCost: 1,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Verify Generic costs are able to be paid when a validation doesn't pass", func() {
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue, Green}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							Blue,
							Green,
							Black,
						},
						GenericCost: 1,
					},
				},
			}
			isMet, combos := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeFalse())
			Expect(combos).To(HaveLen(1))
			Expect(combos[0].GenericCost).To(Equal(0))
			Expect(combos[0].ColorRequirements).To(HaveLen(1))
			Expect(combos[0].ColorRequirements[0]).To(Equal(Black))
		})

		It("All generic objectives can be met with equal number of lands", func() {
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue, Green}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{},
						GenericCost:       4,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})
	})

	When("Playing a land", func() {
		var hand Deck
		var obj TestObjective

		BeforeEach(func() {
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			boardState.Lands = append(boardState.Lands, *CreateUntappedLand([]ManaColor{Red, Blue}))

			hand = NewDeck()
			hand.Cards = []Card{
				{
					Land:    CreateUntappedLand([]ManaColor{White}),
					NonLand: nil,
				},
				{
					Land:    CreateUntappedLand([]ManaColor{Blue}),
					NonLand: nil,
				},
				{
					Land:    CreateUntappedLand([]ManaColor{Red}),
					NonLand: nil,
				},
			}

			obj = TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							White,
							White,
							Blue,
						},
						GenericCost: 0,
					},
				},
			}
		})

		It("Updates the hand and the board correctly", func() {
			updatedHand := boardState.PlayLand(hand, obj, 2)
			Expect(updatedHand.Cards).To(HaveLen(2))
			Expect(boardState.Lands).To(HaveLen(4))
		})
	})

	When("Selecting a land", func() {
		var lands []Card

		BeforeEach(func() {
			lands = []Card{
				*CreateUntappedLandCard([]ManaColor{White}),
				*CreateTappedLandCard([]ManaColor{White, Blue}),
			}
		})

		It("Chooses a land with the highest score", func() {
			costs := []ManaCost{
				{
					ColorRequirements: []ManaColor{White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Blue},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Green},
					GenericCost:       0,
				},
			}

			idx, land := boardState.selectLand(lands, costs, false)
			Expect(idx).To(Equal(1))
			Expect(land).To(Equal(lands[idx]))
		})

		It("Prioritizes an untapped land with the highest score", func() {
			costs := []ManaCost{
				{
					ColorRequirements: []ManaColor{White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Blue},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Green},
					GenericCost:       0,
				},
			}

			idx, land := boardState.selectLand(lands, costs, true)
			Expect(idx).To(Equal(0))
			Expect(land).To(Equal(lands[idx]))
		})

		It("Selects a land in a more complex scenario", func() {
			lands = []Card{
				*CreateUntappedLandCard([]ManaColor{White}),
				*CreateTappedLandCard([]ManaColor{Black, Red}),
				*CreateTappedLandCard([]ManaColor{Green, Red}),
				*CreateTappedLandCard([]ManaColor{White, Blue, Green}),
				*CreateTappedLandCard([]ManaColor{White, Black, Red}),
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White, Red, Black},
					GenericCost:       1,
				},
			}

			idx, land := boardState.selectLand(lands, costOptions, false)
			Expect(idx).To(Equal(4))
			Expect(land).To(Equal(lands[idx]))
		})

		It("Selects a land in a more complex scenario where untapped is priority", func() {
			lands = []Card{
				*CreateUntappedLandCard([]ManaColor{White}),
				*CreateUntappedLandCard([]ManaColor{Black, Red}),
				*CreateTappedLandCard([]ManaColor{Green, Red}),
				*CreateTappedLandCard([]ManaColor{White, Blue, Green}),
				*CreateTappedLandCard([]ManaColor{White, Black, Red}),
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White, Red, Black},
					GenericCost:       1,
				},
			}

			idx, land := boardState.selectLand(lands, costOptions, true)
			Expect(idx).To(Equal(1))
			Expect(land).To(Equal(lands[idx]))
		})
	})

	When("Scoring a land", func() {
		It("computes a simple perfect score correctly", func() {
			land := Land{
				Colors:         []ManaColor{White, Blue},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(12))
		})

		It("computes a perfect score with multiple mana costs correctly", func() {
			land := Land{
				Colors:         []ManaColor{White, Green, Blue},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Green},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Blue},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(33))
		})

		It("computes a simple 0 score correctly", func() {
			land := Land{
				Colors:         []ManaColor{Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(1))
		})

		It("computes a score == colors produced with multiple mana costs correctly", func() {
			land := Land{
				Colors:         []ManaColor{Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Green},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Blue},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(1))
		})

		It("Computes a partially matching score correctly", func() {
			land := Land{
				Colors:         []ManaColor{White, Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White, Blue},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Green, Black},
					GenericCost:       0,
				},
				{
					ColorRequirements: []ManaColor{Blue, Green},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(22))
		})

		It("Doesn't double-count duplicate colors in costs", func() {
			land := Land{
				Colors:         []ManaColor{White, Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []ManaCost{
				{
					ColorRequirements: []ManaColor{White, White},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(12))
		})
	})
})
