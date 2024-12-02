package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
)

var _ = Describe("ObjectiveTranslator", func() {
	Describe("ToInternalTestObjective", func() {
		It("should translate an external API objective to an internal model objective", func() {
			externalObjective := &api.Objective{
				TargetTurn: 1,
				ManaCosts: []*api.ManaCost{
					{
						ColorRequirements: []api.ManaColor{
							api.ManaColor_WHITE,
						},
						GenericCost: 1,
					},
				},
			}
			internalObjective := ToInternalTestObjective(externalObjective)
			Expect(internalObjective.TargetTurn).To(Equal(1))
			Expect(internalObjective.ManaCosts).To(HaveLen(1))
			Expect(internalObjective.ManaCosts[0].ColorRequirements).To(HaveLen(1))
			Expect(internalObjective.ManaCosts[0].GenericCost).To(Equal(1))
		})
	})
})
