package validation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/simulation"
	"manabase-simulation/package/util/test"
)

const (
	SampleDecklist = `
Deck
5 Mountain
4 Manifold Mouse
4 Heartfire Hero
2 Screaming Nemesis
4 Emberheart Challenger
4 Monstrous Rage
2 Rockface Village
2 Obliterating Bolt
4 Thornspire Verge
2 Snakeskin Veil
3 Monastery Swiftspear
4 Burst Lightning
3 Soulstone Sanctuary
3 Questing Druid // Seek the Beast
4 Copperline Gorge
4 Hired Claw
2 Scorching Dragonfire
4 Karplusan Forest

Sideboard
3 Pawpatch Formation
2 Torch the Tower
2 Scorching Shot
2 Tectonic Hazard
1 Questing Druid
2 Lithomantic Barrage
3 Urabrask's Forge

`
	SampleDecklistWithInvalidCard = `
Deck
5 Mountain
3 Soulstone Sanctuary
3 Questing Druid

Sideboard
3 Pawpatch Formation

`
)

var _ = Describe("DefaultParser", func() {
	var parser *DefaultParser

	BeforeEach(func() {
		parser = NewDefaultParser(test.GetDBConfig())
	})

	When("Parsing a decklist", func() {
		It("Correctly parses a decklist", func() {
			deckList, invalidCards, err := parser.Parse(SampleDecklist)
			Expect(err).ToNot(HaveOccurred())
			Expect(deckList.Cards).To(HaveLen(18))
			Expect(simulation.GetTotalCardCount(deckList)).To(Equal(60))
			Expect(invalidCards).To(HaveLen(0))
		})

		It("Correctly identifies invalid cards", func() {
			deckList, invalidCards, err := parser.Parse(SampleDecklistWithInvalidCard)
			Expect(err).ToNot(HaveOccurred())
			Expect(deckList.Cards).To(HaveLen(2))
			Expect(simulation.GetTotalCardCount(deckList)).To(Equal(8))
			Expect(invalidCards).To(HaveLen(1))
			Expect(invalidCards[0].Name).To(Equal("Questing Druid"))
			Expect(invalidCards[0].Reason).To(Equal(MissingCardReason))
		})
	})

	When("Separating a deck and sideboard", func() {
		It("Correctly separates a deck and sideboard", func() {
			deckLines, sideboardLines, err := parser.SeparateDeckAndSideboard([]string{`Deck`, `4 Arclight Phoenix`, `2 Artist's Talent`, ``, `Sideboard`, `2 Abrade`, `1 Anger of the Gods`})
			Expect(err).ToNot(HaveOccurred())
			Expect(deckLines).To(HaveLen(2))
			Expect(sideboardLines).To(HaveLen(2))
			Expect(deckLines[0]).To(Equal("4 Arclight Phoenix"))
			Expect(deckLines[1]).To(Equal("2 Artist's Talent"))
			Expect(sideboardLines[0]).To(Equal("2 Abrade"))
			Expect(sideboardLines[1]).To(Equal("1 Anger of the Gods"))
		})

		It("Correctly handles when sideboard is absent", func() {
			deckLines, sideboardLines, err := parser.SeparateDeckAndSideboard([]string{`Deck`, `4 Arclight Phoenix`, `2 Artist's Talent`, ``})
			Expect(err).ToNot(HaveOccurred())
			Expect(deckLines).To(HaveLen(2))
			Expect(sideboardLines).To(HaveLen(0))
			Expect(deckLines[0]).To(Equal("4 Arclight Phoenix"))
			Expect(deckLines[1]).To(Equal("2 Artist's Talent"))
			Expect(sideboardLines).To(HaveLen(0))
		})

		It("Correctly handles when deck is absent", func() {
			deckLines, sideboardLines, err := parser.SeparateDeckAndSideboard([]string{`Sideboard`, `2 Abrade`, `1 Anger of the Gods`})
			Expect(err).ToNot(HaveOccurred())
			Expect(deckLines).To(HaveLen(0))
			Expect(sideboardLines).To(HaveLen(2))
			Expect(sideboardLines[0]).To(Equal("2 Abrade"))
			Expect(sideboardLines[1]).To(Equal("1 Anger of the Gods"))
		})
	})

	When("Getting names and quantities", func() {
		It("Correctly parses a decklist", func() {
			names, err := parser.getNamesAndQuantities([]string{`1 Arclight Phoenix`, `2 Artist's Talent`, `1 Brazen Borrower`, `4 Consider`})
			Expect(err).ToNot(HaveOccurred())
			Expect(names).To(HaveLen(4))
			Expect(names["Arclight Phoenix"]).To(Equal(1))
			Expect(names["Artist's Talent"]).To(Equal(2))
			Expect(names["Brazen Borrower"]).To(Equal(1))
			Expect(names["Consider"]).To(Equal(4))
		})

		It("Deafults a line with no quantity to 1", func() {
			names, err := parser.getNamesAndQuantities([]string{`Arclight Phoenix`})
			Expect(err).ToNot(HaveOccurred())
			Expect(names).To(HaveLen(1))
			Expect(names["Arclight Phoenix"]).To(Equal(1))
		})
	})
})
