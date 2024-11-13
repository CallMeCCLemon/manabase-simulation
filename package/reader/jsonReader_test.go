package reader

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
)

var _ = Describe("Decklist JSON Parsing Functions", func() {
	When("Reading a JSON Decklist File", func() {
		deck, err := ReadJSONFile[model.DeckList]("../../fixtures/sample_deck.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all of the lands", func() {
			Expect(deck.Lands).To(HaveLen(5))
		})
	})
})
