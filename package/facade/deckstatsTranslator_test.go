package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
)

var _ = Describe("DeckStatsTranslator", func() {
	When("Translating a deck stats", func() {
		It("Should translate a deck stats to an external API definition", func() {
			deckStats := model.DeckStats{
				TotalCards: 10,
				Lands:      5,
				NonLands:   3,
				TotalManaPips: model.SimplifiedManaCost{
					WhiteMana:   1,
					BlueMana:    2,
					GenericMana: 3,
				},
			}
			externalDeckStats := ToExternalDeckStats(deckStats)
			Expect(externalDeckStats.TotalCards).To(Equal(int32(10)))
			Expect(externalDeckStats.Lands).To(Equal(int32(5)))
			Expect(externalDeckStats.NonLands).To(Equal(int32(3)))
			Expect(externalDeckStats.TotalManaPips.WhiteMana).To(Equal(int32(1)))
			Expect(externalDeckStats.TotalManaPips.BlueMana).To(Equal(int32(2)))
			Expect(externalDeckStats.TotalManaPips.GenericCost).To(Equal(int32(3)))
		})
	})
})
