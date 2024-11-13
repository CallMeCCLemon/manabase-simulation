package reader

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScryfallDbReader", func() {
	When("Reading a JSON File", func() {
		cards, err := ReadScryfallDataJSONFile("../../data/scryfall-db.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})

		It("Correctly parses all of the cards", func() {
			Expect(cards).To(HaveLen(33194))
		})
	})
})
