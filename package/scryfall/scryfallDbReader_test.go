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
		_, err := ReadScryfallDataJSONFile("../../data/scryfall-db.json")

		It("Doesn't throw an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("ScryfallDbReader Can Parse and write all of the lands and non-lands in Standard", func() {
	When("Parsing lands cards", func() {
		XIt("Can parse all of the provided land cards in one go and write them to the DB", func() {
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
			Expect(card.Land.Types).To(HaveLen(0))
			Expect(card.Land.Colors).To(HaveLen(1))
			Expect(card.Land.Colors).To(ConsistOf(model.Green))
			Expect(card.Land.EntersTapped).To(BeTrue())
			Expect(card.Land.UntappedCondition.Type).To(Equal(model.ArgothLand))
		})
	})

	When("Parsing non-lands cards", func() {
		XIt("Can parse all of the provided non-land cards in one go and write them to the DB", func() {
			cfg := postgres.Config{
				DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "app", os.Getenv("PORT")),
			}

			db, err := validation.NewPsqlDbAccessor(cfg)
			Expect(err).ToNot(HaveOccurred())
			Expect(db).ToNot(BeNil())

			nonLands, err := reader.ReadJSONFile[[]ScryfallCard]("../../data/non-lands.json")
			Expect(err).ToNot(HaveOccurred())
			parsedNonLands, err := WriteNonLandsToDB(db, nonLands, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(parsedNonLands).To(HaveLen(3080))
		})

		It("Can correctly parse a casting cost", func() {
			cost := `{5}{W}{U}{B}{R}{G}{C}`
			manacost := parseNonLandCastingCost(cost)
			Expect(manacost.ColorRequirements).To(HaveLen(6))
			Expect(manacost.ColorRequirements).To(ConsistOf(model.White, model.Blue, model.Red, model.Black, model.Green, model.Colorless))
			Expect(manacost.GenericCost).To(Equal(5))
		})

		It("Can correctly parse a casting cost", func() {
			cost := `{0}`
			manacost := parseNonLandCastingCost(cost)
			Expect(manacost.ColorRequirements).To(HaveLen(0))
			Expect(manacost.GenericCost).To(Equal(0))
		})

		It("Can correclty parse color pairs", func() {
			cost := `{W/U}{W/B}{W/R}{W/G}{U/B}{U/R}{U/G}{B/R}{B/G}{R/G}`
			manacost := parseNonLandCastingCost(cost)
			Expect(manacost.ColorRequirements).To(HaveLen(10))
			Expect(manacost.ColorRequirements).To(ConsistOf(model.Azorius, model.Orzhov, model.Boros, model.Selesnya, model.Dimir, model.Izzet, model.Simic, model.Rakdos, model.Golgari, model.Gruul))
			Expect(manacost.GenericCost).To(Equal(0))
		})
	})
})
