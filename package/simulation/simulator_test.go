package simulation

import (
	"context"
	"manabase-simulation/package/model"
	"manabase-simulation/package/reader"
	"manabase-simulation/package/util/test"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeckSimulation", func() {
	var deck model.DeckList
	var gameConfig model.GameConfiguration
	var objective model.TestObjective
	var ctx context.Context

	BeforeEach(func() {
		deck, _ = reader.ReadJSONFile[model.DeckList]("../../fixtures/sample_deck.json")
		gameConfig, _ = reader.ReadJSONFile[model.GameConfiguration]("../../fixtures/default-game-config.json")
		objective = model.TestObjective{
			TargetTurn: 3,
			ManaCosts: []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{},
					GenericCost:       1,
				},
			},
		}
		ctx = context.TODO()
	})

	When("Simulating a deck", func() {
		It("Logs correctly to stdout", func() {
			result := SimulateDeck(ctx, deck, gameConfig, objective)
			Expect(result).To(BeTrue())
		})
	})

	When("Simulating Lotus Field", func() {
		deck, err := reader.ReadJSONFile[model.DeckList]("../../fixtures/lotus-field-deck.json")
		Expect(err).ToNot(HaveOccurred())
		gameConfig, _ := reader.ReadJSONFile[model.GameConfiguration]("../../fixtures/default-game-config.json")
		objective := model.TestObjective{
			TargetTurn: 3,
			ManaCosts: []model.ManaCost{
				{
					ColorRequirements: []model.ManaColor{model.White, model.White},
					GenericCost:       1,
				},
			},
		}
		It("Eventually Lotus field succeeds", func() {
			Eventually(func() bool {
				return SimulateDeck(ctx, deck, gameConfig, objective)
			}, 3).Should(BeTrue())
		})
	})
})

var _ = Describe("DeckList", func() {
	When("Getting the card count", func() {
		It("correctly counts all of the lands and nonlands", func() {
			deckList := &model.DeckList{
				Cards: []model.Card{
					*test.CreateUntappedLandCard([]model.ManaColor{model.White}),
					*test.CreateUntappedLandCard([]model.ManaColor{model.White, model.Blue}),
					*test.CreateUntappedLandCard([]model.ManaColor{model.Red, model.Blue}),
					*test.CreateUntappedLandCard([]model.ManaColor{model.Red, model.Blue, model.White}),
					*test.CreateSampleNonLandCard(),
					*test.CreateSampleNonLandCard(),
					*test.CreateSampleNonLandCard(),
				},
			}

			Expect(GetTotalCardCount(deckList)).To(Equal(7))
		})
	})
})
