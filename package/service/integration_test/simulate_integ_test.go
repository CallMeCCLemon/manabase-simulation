package integration_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
)

var _ = Describe("SimulateDeck Integration Test", Label(IntegrationTestLabel), func() {

	It("Simulates a deck", func() {
		resp, err := Client.SimulateDeck(context.Background(), &api.SimulateDeckRequest{
			DeckList: validDeckList,
			GameConfiguration: &api.GameConfiguration{
				InitialHandSize:   7,
				CardsDrawnPerTurn: 1,
				OnThePlay:         true,
			},
			Objective: &api.Objective{
				TargetTurn: 3,
				ManaCosts: []*api.ManaCost{
					{
						ColorRequirements: []api.ManaColor{
							api.ManaColor_RED,
							api.ManaColor_RED,
						},
						GenericCost: 1,
					},
				},
			},
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.Message).To(Equal("The server did the thing!"))
		Expect(resp.SuccessRate > float32(0.1)).To(BeTrue())
		Expect(resp.DeckStats.TotalCards).To(Equal(int32(60)))
		Expect(resp.DeckStats.Lands).To(Equal(int32(22)))
		Expect(resp.DeckStats.NonLands).To(Equal(int32(38)))
		Expect(resp.DeckStats.TotalManaPips.WhiteMana).To(Equal(int32(0)))
		Expect(resp.DeckStats.TotalManaPips.BlueMana).To(Equal(int32(0)))
		Expect(resp.DeckStats.TotalManaPips.BlackMana).To(Equal(int32(0)))
		Expect(resp.DeckStats.TotalManaPips.RedMana).To(Equal(int32(36)))
		Expect(resp.DeckStats.TotalManaPips.GreenMana).To(Equal(int32(5)))
		Expect(resp.DeckStats.TotalManaPips.ColorlessMana).To(Equal(int32(0)))
		Expect(resp.DeckStats.TotalManaPips.GenericCost).To(Equal(int32(19)))

		Expect(resp.DeckStats.LandStats.LandManaProduction.WhiteMana).To(Equal(int32(0)))
		Expect(resp.DeckStats.LandStats.LandManaProduction.BlueMana).To(Equal(int32(0)))
		Expect(resp.DeckStats.LandStats.LandManaProduction.BlackMana).To(Equal(int32(0)))
		Expect(resp.DeckStats.LandStats.LandManaProduction.RedMana).To(Equal(int32(19)))
		Expect(resp.DeckStats.LandStats.LandManaProduction.GreenMana).To(Equal(int32(12)))
		Expect(resp.DeckStats.LandStats.LandManaProduction.ColorlessMana).To(Equal(int32(9)))
		Expect(resp.DeckStats.LandStats.LandManaProduction.GenericCost).To(Equal(int32(0)))
	})

	//It("Returns an error when given an invalid deck", func() {
	//	resp, err := Client.SimulateDeck(context.Background(), &api.SimulateDeckRequest{
	//		DeckList: sampleDecklistWithInvalidCard,
	//		GameConfiguration: &api.GameConfiguration{
	//			InitialHandSize:   7,
	//			CardsDrawnPerTurn: 1,
	//			OnThePlay:         true,
	//		},
	//		Objective: &api.Objective{
	//			TargetTurn: 3,
	//			ManaCosts: []*api.ManaCost{
	//				{
	//					ColorRequirements: []api.ManaColor{
	//						api.ManaColor_WHITE,
	//						api.ManaColor_WHITE,
	//						api.ManaColor_BLUE,
	//					},
	//					GenericCost: 1,
	//				},
	//			},
	//		},
	//	})
	//	Expect(err).ToNot(HaveOccurred())
	//	Expect(resp.Message).To(Equal("The server did the thing!"))
	//	Expect(resp.SuccessRate).To(Equal(float32(0.0)))
	//	Expect(resp.Checkpoints).To(HaveLen(1))
	//	Expect(resp.Checkpoints[0].Iterations).To(Equal(int32(30)))
	//	Expect(resp.Checkpoints[0].Successes).To(Equal(int32(0)))
	//})
})
