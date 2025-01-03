package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

var _ = Describe("ManaCostTranslator", func() {
	When("Translating a mana cost", func() {
		It("Should translate a mana cost to an external API definition", func() {
			manaCost := model.ManaCost{
				ColorRequirements: []model.ManaColor{model.White, model.Blue},
				GenericCost:       1,
			}
			externalManaCost := ToExternalManaCost(&manaCost)
			Expect(externalManaCost.ColorRequirements).To(HaveLen(2))
			Expect(externalManaCost.ColorRequirements[0]).To(Equal(api.ManaColor_WHITE))
			Expect(externalManaCost.ColorRequirements[1]).To(Equal(api.ManaColor_BLUE))
			Expect(externalManaCost.GenericCost).To(Equal(int32(1)))
		})

		It("Should return nil if the mana cost is nil", func() {
			Expect(ToExternalManaCost(nil)).To(BeNil())
		})

		It("Should sensible api mana cost if the mana cost is empty", func() {
			manaCost := ToExternalManaCost(&model.ManaCost{})
			Expect(manaCost).ToNot(BeNil())
			Expect(manaCost.ColorRequirements).To(HaveLen(0))
			Expect(manaCost.GenericCost).To(Equal(int32(0)))
		})
	})
})
