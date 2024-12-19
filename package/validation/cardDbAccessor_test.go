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
		It("Can Read and write a Non-land card", func() {
			card := test.NewLandCard("test-land")
			rowsAffected, err := db.WriteCard(card)
			Expect(err).ToNot(HaveOccurred())
			Expect(rowsAffected).To(Equal(int64(1)))

			newCard, err := db.GetCard(card.Name)
			Expect(err).ToNot(HaveOccurred())
			Expect(newCard).ToNot(BeNil())
			Expect(newCard.NonLand).To(Equal(card.NonLand))
			Expect(newCard.Land).To(Equal(card.Land))
		})

		It("Can Read and write a Land card", func() {
			card := test.NewLandCard("test-land")
			rowsAffected, err := db.WriteCard(card)
			Expect(err).ToNot(HaveOccurred())
			Expect(rowsAffected).To(Equal(int64(1)))

			newCard, err := db.GetCard(card.Name)
			Expect(err).ToNot(HaveOccurred())
			Expect(newCard).ToNot(BeNil())
			Expect(newCard.Land).To(Equal(card.Land))
			Expect(newCard.NonLand).To(Equal(card.NonLand))
		})

		It("Can Do batch reads and writes of lands and non-lands", func() {
			count := 100
			cards := make([]*model.Card, 2*count)
			for i := 0; i < count; i++ {
				cards[i] = test.NewLandCard(fmt.Sprintf("test-land-%d", i))
			}
			for i := 0; i < count; i++ {
				cards[i+count] = test.NewNonLandCard(fmt.Sprintf("test-non-land-%d", i))
			}
			row, err := db.WriteCards(cards)
			Expect(err).ToNot(HaveOccurred())
			Expect(row).To(Equal(int64(2 * count)))
		})
	})
})
