package validation

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"manabase-simulation/package/model"
	"manabase-simulation/package/util/test"
	"os"
)

var _ = Describe("CardDbAccessor", func() {
	var db *CardDbAccessorImpl
	var err error

	BeforeEach(func() {
		cfg := postgres.Config{
			DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "app", os.Getenv("PORT")),
		}

		db, err = NewPsqlDbAccessor(cfg)
		Expect(err).ToNot(HaveOccurred())
		Expect(db).ToNot(BeNil())

		err = db.CreateTables()
		Expect(err).ToNot(HaveOccurred())
	})

	It("Sets up the tables", func() {
		err = db.CreateTables()
		Expect(err).ToNot(HaveOccurred())
	})

	When("Working with the db accessor", func() {
		It("Can create a card", func() {
			card := test.NewLandCard("test-land")
			rowsAffected, err := db.WriteCard(card)
			Expect(err).ToNot(HaveOccurred())
			Expect(rowsAffected).To(Equal(int64(1)))
		})

		It("Can create multiple cards", func() {
			count := 1000
			cards := make([]model.Card, count)
			for i := 0; i < count; i++ {
				cards[i] = *test.NewLandCard(fmt.Sprintf("test-land-%d", i))
			}
			rows, errs := db.WriteCards(cards)
			Expect(errs).ToNot(HaveOccurred())
			Expect(rows).To(Equal(int64(count)))
		})

		It("Can get a card", func() {
			card := test.NewLandCard("test-land")
			rowsAffected, err := db.WriteCard(card)
			Expect(err).ToNot(HaveOccurred())
			Expect(rowsAffected).To(Equal(int64(1)))

			newCard, err := db.GetCard(card.Name)
			Expect(err).ToNot(HaveOccurred())
			Expect(newCard).ToNot(BeNil())
			Expect(newCard.Land.Name).To(Equal(card.Name))
			Expect(newCard.Land.Colors).To(Equal(card.Land.Colors))
			Expect(newCard.Land.EntersTapped).To(Equal(card.Land.EntersTapped))
			Expect(newCard.Land.Types).To(Equal(card.Land.Types))
			Expect(newCard.Land.UntappedCondition).To(Equal(card.Land.UntappedCondition))
			Expect(newCard.Land.ActivationCost).To(Equal(card.Land.ActivationCost))
		})

		//It("Can get a card", func() {
		//	card := model.NewCard(&model.Land{}, nil)
		//	err = db.WriteCard(card)
		//	Expect(err).ToNot(HaveOccurred())
		//
		//	card, err = db.GetCard(card.Land.Name)
		//	Expect(err).ToNot(HaveOccurred())
		//	Expect(card).ToNot(BeNil())
		//	Expect(card.Land.Name).To(Equal(card.Land.Name))
		//})
	})

})

var _ = Describe("toGormModel", func() {
	It("Correctly translates a land to a gorm model", func() {
		card := test.NewLandCard("test-land")
		land, err := toGormModel(card)

		newCard, err := land.Get()
		Expect(err).ToNot(HaveOccurred())
		Expect(newCard.Name).To(Equal(card.Name))
		Expect(newCard.Colors).To(Equal(card.Land.Colors))
		Expect(newCard.EntersTapped).To(Equal(card.Land.EntersTapped))
		Expect(newCard.Types).To(Equal(card.Land.Types))
		Expect(newCard.UntappedCondition).To(Equal(card.Land.UntappedCondition))
		Expect(newCard.ActivationCost).To(Equal(card.Land.ActivationCost))
	})
})
