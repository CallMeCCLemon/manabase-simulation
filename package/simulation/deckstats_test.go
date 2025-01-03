package simulation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
	})
})
