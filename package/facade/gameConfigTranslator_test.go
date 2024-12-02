package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
)

var _ = Describe("GameConfigTranslator", func() {
	When("translating a game configuration", func() {
		It("should translate the game configuration to a game config", func() {
			gameConfig := api.GameConfiguration{
				InitialHandSize:   7,
				CardsDrawnPerTurn: 1,
				OnThePlay:         true,
			}
			gameConfigTranslator := ToInternalGameConfiguration(&gameConfig)

			Expect(gameConfigTranslator.InitialHandSize).To(Equal(7))
			Expect(gameConfigTranslator.CardsDrawnPerTurn).To(Equal(1))
			Expect(gameConfigTranslator.OnThePlay).To(Equal(true))
		})
	})
})
