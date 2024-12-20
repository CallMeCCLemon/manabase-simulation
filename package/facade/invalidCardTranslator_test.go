package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
)

var _ = Describe("InvalidCardTranslator", func() {
	When("Converting an invalid card", func() {
		It("Correctly converts an invalid card", func() {
			invalidCard := model.InvalidCard{
				Name:   "test-card",
				Reason: "test-reason",
			}
			externalInvalidCard := ToExternalInvalidCard(invalidCard)
			Expect(externalInvalidCard.Name).To(Equal(invalidCard.Name))
			Expect(externalInvalidCard.Reason).To(Equal(invalidCard.Reason))

			internalInvalidCard := ToInternalInvalidCard(externalInvalidCard)
			Expect(internalInvalidCard.Name).To(Equal(invalidCard.Name))
			Expect(internalInvalidCard.Reason).To(Equal(invalidCard.Reason))
		})
	})
})
