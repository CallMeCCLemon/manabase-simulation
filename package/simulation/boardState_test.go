package simulation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
	"manabase-simulation/package/util/test"
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
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Blue}))

			obj := model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.White,
							model.White,
							model.Blue,
						},
						GenericCost: 0,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Is able to solve for an objective with basic and dual lands as met", func() {
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue}))

			obj := model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.White,
							model.White,
							model.Blue,
						},
						GenericCost: 0,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())

			obj = model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.White,
							model.Red,
							model.Blue,
						},
						GenericCost: 0,
					},
				},
			}

			isMet, _ = boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Is able to solve for an objective with basic, dual, and triome lands as met", func() {
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue, model.Green}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue}))

			obj := model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.White,
							model.White,
							model.Blue,
							model.Red,
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
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue, model.Green}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue}))

			obj := model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.Blue,
							model.Blue,
							model.Blue,
							model.Green,
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
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue, model.Green}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue}))

			obj := model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.Blue,
							model.Blue,
							model.Green,
						},
						GenericCost: 1,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})

		It("Verify Generic costs are able to be paid when a validation doesn't pass", func() {
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue, model.Green}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue}))

			obj := model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.Blue,
							model.Green,
							model.Black,
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
			Expect(combos[0].ColorRequirements[0]).To(Equal(model.Black))
		})

		It("All generic objectives can be met with equal number of lands", func() {
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue, model.Green}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue}))

			obj := model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{},
						GenericCost:       4,
					},
				},
			}
			isMet, _ := boardState.ValidateTestObjective(obj)
			Expect(isMet).To(BeTrue())
		})
	})

	When("Playing a land", func() {
		var hand model.Deck
		var obj model.TestObjective

		BeforeEach(func() {
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			boardState.Lands = append(boardState.Lands, *test.CreateUntappedLand([]model.ManaColor{model.Red, model.Blue}))

			hand = model.NewDeck()
			hand.Cards = []model.Card{
				{
					Land:    test.CreateUntappedLand([]model.ManaColor{model.White}),
					NonLand: nil,
				},
				{
					Land:    test.CreateUntappedLand([]model.ManaColor{model.Blue}),
					NonLand: nil,
				},
				{
					Land:    test.CreateUntappedLand([]model.ManaColor{model.Red}),
					NonLand: nil,
				},
			}

			obj = model.TestObjective{
				TargetTurn: 3,
				ManaCosts: []model.ManaCost{
					{
						ColorRequirements: []model.ManaColor{
							model.White,
							model.White,
							model.Blue,
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
		var lands []model.Card

		BeforeEach(func() {
			lands = []model.Card{
				*test.CreateUntappedLandCard([]model.ManaColor{model.White}),
				*test.CreateTappedLandCard([]model.ManaColor{model.White, model.Blue}),
			}
		})

		It("Chooses a land with the highest score", func() {
			costs := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Blue},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Green},
					GenericCost:       0,
				},
			}

			idx, land := boardState.selectLand(lands, costs, false)
			Expect(idx).To(Equal(1))
			Expect(land).To(Equal(lands[idx]))
		})

		It("Prioritizes an untapped land with the highest score", func() {
			costs := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Blue},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Green},
					GenericCost:       0,
				},
			}

			idx, land := boardState.selectLand(lands, costs, true)
			Expect(idx).To(Equal(0))
			Expect(land).To(Equal(lands[idx]))
		})

		It("Selects a land in a more complex scenario", func() {
			lands = []model.Card{
				*test.CreateUntappedLandCard([]model.ManaColor{model.White}),
				*test.CreateTappedLandCard([]model.ManaColor{model.Black, model.Red}),
				*test.CreateTappedLandCard([]model.ManaColor{model.Green, model.Red}),
				*test.CreateTappedLandCard([]model.ManaColor{model.White, model.Blue, model.Green}),
				*test.CreateTappedLandCard([]model.ManaColor{model.White, model.Black, model.Red}),
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White, model.Red, model.Black},
					GenericCost:       1,
				},
			}

			idx, land := boardState.selectLand(lands, costOptions, false)
			Expect(idx).To(Equal(4))
			Expect(land).To(Equal(lands[idx]))
		})

		It("Selects a land in a more complex scenario where untapped is priority", func() {
			lands = []model.Card{
				*test.CreateUntappedLandCard([]model.ManaColor{model.White}),
				*test.CreateUntappedLandCard([]model.ManaColor{model.Black, model.Red}),
				*test.CreateTappedLandCard([]model.ManaColor{model.Green, model.Red}),
				*test.CreateTappedLandCard([]model.ManaColor{model.White, model.Blue, model.Green}),
				*test.CreateTappedLandCard([]model.ManaColor{model.White, model.Black, model.Red}),
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White, model.Red, model.Black},
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
			land := model.Land{
				Colors:         []model.ManaColor{model.White, model.Blue},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(12))
		})

		It("computes a perfect score with multiple mana costs correctly", func() {
			land := model.Land{
				Colors:         []model.ManaColor{model.White, model.Green, model.Blue},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Green},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Blue},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(33))
		})

		It("computes a simple 0 score correctly", func() {
			land := model.Land{
				Colors:         []model.ManaColor{model.Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(1))
		})

		It("computes a score == colors produced with multiple mana costs correctly", func() {
			land := model.Land{
				Colors:         []model.ManaColor{model.Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Green},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Blue},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(1))
		})

		It("Computes a partially matching score correctly", func() {
			land := model.Land{
				Colors:         []model.ManaColor{model.White, model.Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White, model.Blue},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Green, model.Black},
					GenericCost:       0,
				},
				{
					ColorRequirements: []model.ManaColor{model.Blue, model.Green},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(22))
		})

		It("Doesn't double-count duplicate colors in costs", func() {
			land := model.Land{
				Colors:         []model.ManaColor{model.White, model.Black},
				EntersTapped:   false,
				ActivationCost: nil,
			}

			costOptions := []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White, model.White},
					GenericCost:       0,
				},
			}

			Expect(scoreLand(land, costOptions)).To(Equal(12))
		})
	})
})

var _ = Describe("ShockLand", func() {
	When("Determining if the land can enter untapped", func() {
		It("Can enter untapped when there is 2 life", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.ShockLand,
				},
			}

			boardState := BoardState{
				Life: 3,
			}
			Expect(boardState.CanEnterUntapped(land)).To(BeTrue())
		})

		It("Can NOT enter untapped when there is not enough life", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.ShockLand,
				},
			}

			boardState := BoardState{
				Life: 1,
			}
			Expect(boardState.CanEnterUntapped(land)).To(BeFalse())
		})
	})

	When("Paying the untapped cost", func() {
		It("Pays the cost correctly", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.ShockLand,
				},
			}

			boardState := BoardState{
				Life: 3,
			}
			Expect(boardState.PayUntappedCost(&land)).To(BeNil())
			Expect(boardState.Life).To(Equal(1))
		})

		It("Doesn't pay the untapped cost when there is not enough life", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.ShockLand,
				},
			}

			boardState := BoardState{
				Life: 2,
			}
			Expect(boardState.PayUntappedCost(&land)).To(HaveOccurred())
			Expect(boardState.Life).To(Equal(2))
		})

		It("Doesn't pay the untapped cost when it's not tapped", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: false,
				UntappedCondition: &model.UntappedCondition{
					Type: model.ShockLand,
				},
			}

			boardState := BoardState{
				Life: 2,
			}
			Expect(boardState.PayUntappedCost(&land)).ToNot(HaveOccurred())
			Expect(boardState.Life).To(Equal(2))
		})
	})
})

var _ = Describe("FastLand", func() {
	When("Determining if the land can enter untapped", func() {
		It("Can enter untapped when there are 2 or less lands", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.FastLand,
				},
			}

			boardState := BoardState{
				Life: 3,
				Lands: []model.Land{
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
				},
			}
			Expect(boardState.CanEnterUntapped(land)).To(BeTrue())
		})

		It("Can NOT enter untapped when there are 3 or more lands", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.FastLand,
				},
			}

			boardState := BoardState{
				Life: 3,
				Lands: []model.Land{
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
				},
			}
			Expect(boardState.CanEnterUntapped(land)).To(BeFalse())
		})
	})

	When("Paying the untapped cost", func() {
		It("Pays the cost correctly", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.FastLand,
				},
			}

			boardState := BoardState{
				Life: 3,
			}
			Expect(boardState.PayUntappedCost(&land)).ToNot(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})

		It("Doesn't pay the untapped cost when there are too many lands", func() {
			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.FastLand,
				},
			}

			boardState := BoardState{
				Life: 3,
				Lands: []model.Land{
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
				},
			}
			Expect(boardState.PayUntappedCost(&land)).To(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})
	})
})

var _ = Describe("CheckLand", func() {
	When("Determining if the land can enter untapped", func() {
		It("Can enter untapped when there is a land of the right type", func() {
			c := model.CheckLandData{model.Plains, model.Forest}
			cString, _ := c.ToString()

			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.CheckLand,
					Data: &cString,
				},
			}

			l := test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue})
			l.Types = []model.LandType{model.Plains}

			boardState := BoardState{
				Life: 3,
				Lands: []model.Land{
					*l,
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
				},
			}
			Expect(boardState.CanEnterUntapped(land)).To(BeTrue())
		})

		It("Can NOT enter untapped when there is no land of the right type", func() {
			c := model.CheckLandData{model.Plains, model.Forest}
			cString, _ := c.ToString()

			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.CheckLand,
					Data: &cString,
				},
			}

			l := test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue})
			l.Types = []model.LandType{model.Swamp}

			boardState := BoardState{
				Life: 3,
				Lands: []model.Land{
					*l,
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
				},
			}
			Expect(boardState.CanEnterUntapped(land)).To(BeFalse())
		})
	})

	When("Paying the untapped cost", func() {
		It("Pays the cost correctly", func() {
			c := model.CheckLandData{model.Plains, model.Forest}
			cString, _ := c.ToString()

			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.CheckLand,
					Data: &cString,
				},
			}

			l := test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue})
			l.Types = []model.LandType{model.Plains}

			boardState := BoardState{
				Life: 3,
				Lands: []model.Land{
					*l,
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
				},
			}
			Expect(boardState.PayUntappedCost(&land)).ToNot(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})

		It("Doesn't pay the untapped cost when there is no land of the right type", func() {
			c := model.CheckLandData{model.Plains, model.Forest}
			cString, _ := c.ToString()

			land := model.Land{
				Colors:       []model.ManaColor{model.White, model.Blue},
				EntersTapped: true,
				UntappedCondition: &model.UntappedCondition{
					Type: model.CheckLand,
					Data: &cString,
				},
			}

			l := test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue})
			l.Types = []model.LandType{model.Swamp}

			boardState := BoardState{
				Life: 3,
				Lands: []model.Land{
					*l,
					*test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}),
				},
			}
			Expect(boardState.PayUntappedCost(&land)).To(HaveOccurred())
			Expect(boardState.Life).To(Equal(3))
		})
	})
})

var _ = Describe("Sorting a list of lands", func() {
	When("Sorting a list of lands with different quantities of mana they can tap for", func() {
		It("Returns a list where each subsequent len of land.Colors >= prevLand.Colors", func() {
			var lands []model.Land
			lands = append(lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Green, model.Red, model.Black}))
			lands = append(lands, *test.CreateUntappedLand([]model.ManaColor{model.White}))
			lands = append(lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Green, model.Red}))
			lands = append(lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Blue}))
			lands = append(lands, *test.CreateUntappedLand([]model.ManaColor{model.White, model.Green, model.Red, model.Black, model.Blue}))

			sortedLands := SortLandsByRestrictiveness(lands)

			Expect(sortedLands[0].Colors).To(HaveLen(1))
			Expect(sortedLands[1].Colors).To(HaveLen(2))
			Expect(sortedLands[2].Colors).To(HaveLen(3))
			Expect(sortedLands[3].Colors).To(HaveLen(4))
			Expect(sortedLands[4].Colors).To(HaveLen(5))
		})
	})
})
