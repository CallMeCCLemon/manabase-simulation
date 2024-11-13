package model

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sorting a list of lands", func() {
	When("Sorting a list of lands with different quantities of mana they can tap for", func() {
		It("Returns a list where each subsequent len of land.Colors >= prevLand.Colors", func() {
			var lands []Land
			lands = append(lands, *CreateUntappedLand([]ManaColor{White, Green, Red, Black}))
			lands = append(lands, *CreateUntappedLand([]ManaColor{White}))
			lands = append(lands, *CreateUntappedLand([]ManaColor{White, Green, Red}))
			lands = append(lands, *CreateUntappedLand([]ManaColor{White, Blue}))
			lands = append(lands, *CreateUntappedLand([]ManaColor{White, Green, Red, Black, Blue}))

			sortedLands := SortLandsByRestrictiveness(lands)

			Expect(sortedLands[0].Colors).To(HaveLen(1))
			Expect(sortedLands[1].Colors).To(HaveLen(2))
			Expect(sortedLands[2].Colors).To(HaveLen(3))
			Expect(sortedLands[3].Colors).To(HaveLen(4))
			Expect(sortedLands[4].Colors).To(HaveLen(5))
		})
	})
})
