package simulation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
	"manabase-simulation/package/util/test"
)

var _ = Describe("DeckStats", func() {
	When("Getting the deck stats", func() {
		It("Should return the correct deck stats", func() {
			deckList := test.NewDeckList()
			stats := GetDeckStats(*deckList)
			Expect(stats.TotalCards).To(Equal(4))
			Expect(stats.Lands).To(Equal(2))
			Expect(stats.NonLands).To(Equal(2))
			Expect(stats.TotalManaPips.WhiteMana).To(Equal(2))
			Expect(stats.TotalManaPips.GenericMana).To(Equal(2))
		})

		It("Can return deck stats for all land cards", func() {
			deckList := model.DeckList{
				Cards: []model.Card{
					*test.CreateUntappedLandCard([]model.ManaColor{model.White}),
					*test.CreateUntappedLandCard([]model.ManaColor{model.White, model.Blue}),
					*test.CreateUntappedLandCard([]model.ManaColor{model.Red, model.Blue, model.White, model.Green, model.Black, model.Colorless}),
					*test.CreateUntappedLandCard([]model.ManaColor{model.Red, model.Blue, model.White}),
				},
			}
			stats := GetDeckStats(deckList)
			Expect(stats.TotalCards).To(Equal(4))
			Expect(stats.Lands).To(Equal(4))
			Expect(stats.NonLands).To(Equal(0))
			Expect(stats.TotalManaPips.WhiteMana).To(Equal(0))
			Expect(stats.TotalManaPips.BlueMana).To(Equal(0))
			Expect(stats.TotalManaPips.BlackMana).To(Equal(0))
			Expect(stats.TotalManaPips.RedMana).To(Equal(0))
			Expect(stats.TotalManaPips.GreenMana).To(Equal(0))
			Expect(stats.TotalManaPips.ColorlessMana).To(Equal(0))
			Expect(stats.LandStats.LandManaProduction.WhiteMana).To(Equal(4))
			Expect(stats.LandStats.LandManaProduction.BlueMana).To(Equal(3))
			Expect(stats.LandStats.LandManaProduction.BlackMana).To(Equal(1))
			Expect(stats.LandStats.LandManaProduction.RedMana).To(Equal(2))
			Expect(stats.LandStats.LandManaProduction.GreenMana).To(Equal(1))
			Expect(stats.LandStats.LandManaProduction.ColorlessMana).To(Equal(1))
		})
	})
})
