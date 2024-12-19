package scryfall

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"manabase-simulation/package/model"
	"manabase-simulation/package/reader"
	"manabase-simulation/package/validation"
	"os"
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

var _ = Describe("ScryfallDbReader Do theThing", func() {
	When("Parsing lands cards", func() {
		It("Can parse all of the provided land cards in one go and write them to the DB", func() {
			cfg := postgres.Config{
				DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "app", os.Getenv("PORT")),
			}

			db, err := validation.NewPsqlDbAccessor(cfg)
			Expect(err).ToNot(HaveOccurred())
			Expect(db).ToNot(BeNil())

			lands, err := reader.ReadJSONFile[[]ScryfallCard]("../../data/lands.json")
			Expect(err).ToNot(HaveOccurred())
			parsedLands, err := WriteLandsToDB(db, lands, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(parsedLands).To(HaveLen(190))
		})

		It("Correctly parses a tapped land", func() {
			land := ScryfallCard{
				Name:         "Test-Land",
				TypeLine:     "Land — Plains",
				ProducedMana: []string{"W"},
				OracleText:   "Test-Land enters tapped",
			}
			card := parseLandCard(land)
			Expect(card.Name).To(Equal(land.Name))
			Expect(card.Land.Name).To(Equal(land.Name))
			Expect(card.Land.Types).To(HaveLen(1))
			Expect(card.Land.Types[0]).To(Equal(model.Plains))
			Expect(card.Land.Colors).To(HaveLen(1))
			Expect(card.Land.Colors[0]).To(Equal(model.White))
			Expect(card.Land.EntersTapped).To(BeTrue())
			Expect(card.Land.UntappedCondition).To(BeNil())
		})

		It("Correctly parses a fast land", func() {
			land := ScryfallCard{
				Name:         "Test-Land",
				TypeLine:     "Land",
				ProducedMana: []string{"U", "G"},
				OracleText:   "Test-Land enters tapped unless you control two or fewer other lands.",
			}
			card := parseLandCard(land)
			Expect(card.Name).To(Equal(land.Name))
			Expect(card.Land.Name).To(Equal(land.Name))
			Expect(card.Land.Types).To(HaveLen(0))
			Expect(card.Land.Colors).To(HaveLen(2))
			Expect(card.Land.Colors).To(ConsistOf(model.Blue, model.Green))
			Expect(card.Land.EntersTapped).To(BeTrue())
			Expect(card.Land.UntappedCondition.Type).To(Equal(model.FastLand))
		})

		It("Correctly parses an unlucky land", func() {
			land := ScryfallCard{
				Name:         "Test-Land",
				TypeLine:     "Land",
				ProducedMana: []string{"B", "R"},
				OracleText:   "Test-Land enters tapped unless a player has 13 or less life.",
			}
			card := parseLandCard(land)
			Expect(card.Name).To(Equal(land.Name))
			Expect(card.Land.Name).To(Equal(land.Name))
			Expect(card.Land.Types).To(HaveLen(0))
			Expect(card.Land.Colors).To(HaveLen(2))
			Expect(card.Land.Colors).To(ConsistOf(model.Black, model.Red))
			Expect(card.Land.EntersTapped).To(BeTrue())
			Expect(card.Land.UntappedCondition.Type).To(Equal(model.UnluckyLand))
		})

		It("Correctly parses a Typal land", func() {
			land := ScryfallCard{
				Name:         "Test-Land",
				TypeLine:     "Land",
				ProducedMana: []string{"W"},
				OracleText:   "Test-Land enters tapped unless you revealed a Soldier card this way or you control a Soldier.",
			}
			card := parseLandCard(land)
			Expect(card.Name).To(Equal(land.Name))
			Expect(card.Land.Name).To(Equal(land.Name))
			Expect(card.Land.Types).To(HaveLen(0))
			Expect(card.Land.Colors).To(HaveLen(1))
			Expect(card.Land.Colors).To(ConsistOf(model.White))
			Expect(card.Land.EntersTapped).To(BeTrue())
			Expect(card.Land.UntappedCondition.Type).To(Equal(model.TypalLand))
		})

		It("Correctly parses a Argoth land", func() {
			land := ScryfallCard{
				Name:         "Test-Land",
				TypeLine:     "Land",
				ProducedMana: []string{"G"},
				OracleText:   "Test-Land enters tapped unless you control a legendary green creature.",
			}
			card := parseLandCard(land)
			Expect(card.Name).To(Equal(land.Name))
			Expect(card.Land.Name).To(Equal(land.Name))
			Expect(card.Land.Types).To(HaveLen(0))
			Expect(card.Land.Colors).To(HaveLen(1))
			Expect(card.Land.Colors).To(ConsistOf(model.Green))
			Expect(card.Land.EntersTapped).To(BeTrue())
			Expect(card.Land.UntappedCondition.Type).To(Equal(model.ArgothLand))
		})
	})
})
