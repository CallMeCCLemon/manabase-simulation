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
	})
	//
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
