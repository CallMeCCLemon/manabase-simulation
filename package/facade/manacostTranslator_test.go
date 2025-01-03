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

	When("Translating a simplified mana cost", func() {
		It("Should translate a simplified mana cost to an external API definition", func() {
			manaCost := model.ManaCost{
				ColorRequirements: []model.ManaColor{model.White, model.White, model.White, model.White, model.Blue, model.Red, model.Green, model.Colorless, model.Black},
				GenericCost:       3,
			}
			externalManaCost := ToExternalSimplifiedManaCost(&manaCost)
			Expect(externalManaCost.WhiteMana).To(Equal(int32(4)))
			Expect(externalManaCost.BlueMana).To(Equal(int32(1)))
			Expect(externalManaCost.BlackMana).To(Equal(int32(1)))
			Expect(externalManaCost.RedMana).To(Equal(int32(1)))
			Expect(externalManaCost.GreenMana).To(Equal(int32(1)))
			Expect(externalManaCost.ColorlessMana).To(Equal(int32(1)))
			Expect(externalManaCost.GenericCost).To(Equal(int32(3)))
		})

		It("Should return nil if the mana cost is nil", func() {
			Expect(ToExternalSimplifiedManaCost(nil)).To(BeNil())
		})

		It("Should sensible api mana cost if the mana cost is empty", func() {
			manaCost := ToExternalSimplifiedManaCost(&model.ManaCost{})
			Expect(manaCost).ToNot(BeNil())
			Expect(manaCost.WhiteMana).To(Equal(int32(0)))
			Expect(manaCost.BlueMana).To(Equal(int32(0)))
			Expect(manaCost.BlackMana).To(Equal(int32(0)))
			Expect(manaCost.RedMana).To(Equal(int32(0)))
			Expect(manaCost.GreenMana).To(Equal(int32(0)))
			Expect(manaCost.ColorlessMana).To(Equal(int32(0)))
			Expect(manaCost.GenericCost).To(Equal(int32(0)))
		})
	})
})
