package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

var _ = Describe("DecklistTranslator", func() {
	Describe("ToInternalDeckList", func() {
		It("should translate an external decklist to the internal model", func() {
			sampleDeckList := &api.DeckList{
				Lands: []*api.Land{
					{
						Name:              "Forest",
						Types:             []api.LandType{api.LandType_FOREST},
						Colors:            []api.ManaColor{api.ManaColor_BLUE},
						EntersTapped:      true,
						ActivationCost:    &api.ActivationCost{Life: 1, ManaCost: &api.ManaCost{GenericCost: 2}},
						UntappedCondition: &api.UntappedCondition{Type: api.ConditionType_SHOCK_LAND, Data: "something-not-null"},
						Quantity:          1,
					},
					{
						Name:              "Plains",
						Types:             []api.LandType{api.LandType_PLAINS},
						Colors:            []api.ManaColor{api.ManaColor_BLUE, api.ManaColor_GREEN},
						EntersTapped:      false,
						ActivationCost:    nil,
						UntappedCondition: nil,
						Quantity:          2,
					},
					{
						Name:              "Mountain",
						Types:             []api.LandType{api.LandType_MOUNTAIN},
						Colors:            []api.ManaColor{api.ManaColor_BLUE, api.ManaColor_GREEN, api.ManaColor_RED},
						EntersTapped:      false,
						ActivationCost:    nil,
						UntappedCondition: nil,
						Quantity:          3,
					},
				},
				NonLands: []*api.NonLand{
					{
						Name:        "Llanowar Elves",
						CastingCost: &api.ManaCost{GenericCost: 1},
						Quantity:    1,
					},
					{
						Name: "Llanowar Goblins",
						CastingCost: &api.ManaCost{
							ColorRequirements: []api.ManaColor{api.ManaColor_BLUE, api.ManaColor_GREEN},
							GenericCost:       2,
						},
						Quantity: 2,
					},
				},
			}
			decklist := ToInternalDeckList(sampleDeckList)

			Expect(decklist.Lands).To(HaveLen(3))
			Expect(decklist.NonLands).To(HaveLen(2))
		})
	})

	Describe("toInternalLand", func() {
		It("should translate an external land to the internal model", func() {
			sampleLand := &api.Land{
				Name:              "Forest",
				Types:             []api.LandType{api.LandType_FOREST},
				Colors:            []api.ManaColor{api.ManaColor_BLUE},
				EntersTapped:      true,
				ActivationCost:    &api.ActivationCost{Life: 1, ManaCost: &api.ManaCost{GenericCost: 2}},
				UntappedCondition: &api.UntappedCondition{Type: api.ConditionType_SHOCK_LAND, Data: "something-not-null"},
				Quantity:          1,
			}
			land := toInternalLand(sampleLand)

			Expect(land.Name).To(Equal("Forest"))
			Expect(land.Types).To(Equal([]model.LandType{model.Forest}))
			Expect(land.Colors).To(Equal([]model.ManaColor{model.Blue}))
			Expect(land.EntersTapped).To(BeTrue())
			Expect(land.ActivationCost).ToNot(BeNil())
			Expect(land.UntappedCondition).ToNot(BeNil())
			Expect(land.Quantity).To(Equal(1))
		})
	})

	Describe("toInternalLandType", func() {
		It("should translate a land type to the internal model", func() {
			Expect(toInternalLandType(api.LandType_FOREST)).To(Equal(model.Forest))
			Expect(toInternalLandType(api.LandType_ISLAND)).To(Equal(model.Island))
			Expect(toInternalLandType(api.LandType_MOUNTAIN)).To(Equal(model.Mountain))
			Expect(toInternalLandType(api.LandType_SWAMP)).To(Equal(model.Swamp))
			Expect(toInternalLandType(api.LandType_PLAINS)).To(Equal(model.Plains))
		})
	})

	Describe("toInternalManaColor", func() {
		It("should translate a mana color to the internal model", func() {
			Expect(toInternalManaColor(api.ManaColor_WHITE)).To(Equal(model.White))
			Expect(toInternalManaColor(api.ManaColor_BLUE)).To(Equal(model.Blue))
			Expect(toInternalManaColor(api.ManaColor_BLACK)).To(Equal(model.Black))
			Expect(toInternalManaColor(api.ManaColor_RED)).To(Equal(model.Red))
			Expect(toInternalManaColor(api.ManaColor_GREEN)).To(Equal(model.Green))
			Expect(toInternalManaColor(api.ManaColor_COLORLESS)).To(Equal(model.Colorless))
		})
	})

	Describe("toInternalNonLand", func() {
		It("should translate an external non land to the internal model", func() {
			sampleNonLand := &api.NonLand{
				Name:        "Llanowar Elves",
				CastingCost: &api.ManaCost{GenericCost: 1},
				Quantity:    1,
			}
			nonLand := toInternalNonLand(sampleNonLand)

			Expect(nonLand.Name).To(Equal("Llanowar Elves"))
			Expect(nonLand.CastingCost).ToNot(BeNil())
			Expect(nonLand.CastingCost.GenericCost).To(Equal(1))
			Expect(nonLand.Quantity).To(Equal(1))
		})
	})

	Describe("toInternalActivationCost", func() {
		It("should translate an activation cost to the internal model", func() {
			sampleActivationCost := &api.ActivationCost{
				ManaCost: &api.ManaCost{
					ColorRequirements: []api.ManaColor{api.ManaColor_WHITE, api.ManaColor_BLUE},
					GenericCost:       1,
				},
				Life: 4,
			}
			activationCost := toInternalActivationCost(sampleActivationCost)

			Expect(*activationCost.ManaCost).To(Equal(model.ManaCost{
				ColorRequirements: []model.ManaColor{model.White, model.Blue},
				GenericCost:       1,
			}))
			Expect(*activationCost.Life).To(Equal(4))
		})

		It("should return nil if the activation cost is nil", func() {
			Expect(toInternalActivationCost(nil)).To(BeNil())
		})
	})

	Describe("toInternalUntappedCondition", func() {
		It("should translate an untapped condition to the internal model", func() {
			sampleUntappedCondition := &api.UntappedCondition{
				Type: api.ConditionType_SHOCK_LAND,
				Data: "something",
			}
			untappedCondition := toInternalUntappedCondition(sampleUntappedCondition)

			Expect(untappedCondition.Type).To(Equal(model.ShockLand))
			Expect(*untappedCondition.Data).To(Equal("something"))
		})

		It("should return nil if the untapped condition is nil", func() {
			Expect(toInternalUntappedCondition(nil)).To(BeNil())
		})
	})

	Describe("toInternalConditionType", func() {
		It("should translate a condition type to the internal model", func() {
			Expect(toInternalConditionType(api.ConditionType_SHOCK_LAND)).To(Equal(model.ShockLand))
			Expect(toInternalConditionType(api.ConditionType_FAST_LAND)).To(Equal(model.FastLand))
			Expect(toInternalConditionType(api.ConditionType_CHECK_LAND)).To(Equal(model.CheckLand))
		})
	})

	Describe("toInternalManaCost", func() {
		It("should translate a mana cost to the internal model", func() {
			sampleManaCost := &api.ManaCost{
				ColorRequirements: []api.ManaColor{api.ManaColor_WHITE, api.ManaColor_BLUE},
				GenericCost:       1,
			}
			manaCost := toInternalManaCost(sampleManaCost)

			Expect(manaCost.ColorRequirements).To(Equal([]model.ManaColor{model.White, model.Blue}))
			Expect(manaCost.GenericCost).To(Equal(1))
		})

		It("Should return nil if the mana cost is nil", func() {
			Expect(toInternalManaCost(nil)).To(BeNil())
		})
	})
})
