package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
	"manabase-simulation/package/model"
)

var _ = Describe("DeckStatsTranslator", func() {
	When("Translating a deck stats", func() {
		It("Should translate a deck stats to an external API definition", func() {
			deckStats := model.DeckStats{
				TotalCards: 10,
				Lands:      5,
				NonLands:   3,
				TotalManaPips: model.ManaCost{
					ColorRequirements: []model.ManaColor{model.White, model.Blue},
					GenericCost:       1,
				},
			}
			externalDeckStats := ToExternalDeckStats(deckStats)
			Expect(externalDeckStats.TotalCards).To(Equal(int32(10)))
			Expect(externalDeckStats.Lands).To(Equal(int32(5)))
			Expect(externalDeckStats.NonLands).To(Equal(int32(3)))
			Expect(externalDeckStats.TotalManaPips.ColorRequirements).To(HaveLen(2))
			Expect(externalDeckStats.TotalManaPips.ColorRequirements[0]).To(Equal(api.ManaColor_WHITE))
			Expect(externalDeckStats.TotalManaPips.ColorRequirements[1]).To(Equal(api.ManaColor_BLUE))
			Expect(externalDeckStats.TotalManaPips.GenericCost).To(Equal(int32(1)))
		})
	})
})
