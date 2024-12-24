package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

var _ = Describe("DecklistTranslator", func() {
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
