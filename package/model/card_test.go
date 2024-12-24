package model

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/reader"
)

var _ = Describe("Card", func() {
	When("Reading a card from a JSON file", func() {
		It("Correctly reads a card from a JSON file", func() {
			card, err := reader.ReadJSONFile[Card]("../../fixtures/sample_card.json")
			Expect(err).ToNot(HaveOccurred())
			Expect(card.Name).To(Equal("Test-Mountain"))
			Expect(card.Land).ToNot(BeNil())
			Expect(card.Land.Colors).To(HaveLen(1))
			Expect(card.Land.Colors[0]).To(Equal(Red))
			Expect(card.Land.EntersTapped).To(BeFalse())
			Expect(card.NonLand).To(BeNil())
		})
	})
})
