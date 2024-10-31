package main

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
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							white,
							white,
							blue,
						},
						GenericCost: 0,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Is able to solve for an objective with basic and dual lands as met", func() {
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white, blue}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							white,
							white,
							blue,
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
							white,
							red,
							blue,
						},
						GenericCost: 0,
					},
				},
			}

			isMet, _ = boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Is able to solve for an objective with basic, dual, and triome lands as met", func() {
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white, blue}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue, green}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							white,
							white,
							blue,
							red,
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
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white, blue}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue, green}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							blue,
							blue,
							blue,
							green,
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
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white, blue}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue, green}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							blue,
							blue,
							green,
						},
						GenericCost: 1,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Verify Generic costs are able to be paid when a validation doesn't pass", func() {
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{white, blue}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue, green}))
			boardState.Lands = append(boardState.Lands, createUntappedLand([]ManaColor{red, blue}))

			obj := TestObjective{
				TargetTurn: 3,
				ManaCosts: []ManaCost{
					{
						ColorRequirements: []ManaColor{
							blue,
							green,
							black,
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
			Expect(combos[0].ColorRequirements[0]).To(Equal(black))
		})
	})
})
