package facade

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/package/model"
)

var _ = Describe("CheckpointTranslator", func() {
	When("Translating a checkpoint", func() {
		It("Should translate a checkpoint to an external API definition", func() {
			checkpoint := model.ResultCheckpoint{
				Iterations: 10,
				Successes:  5,
			}
			externalCheckpoint := ToExternalResultCheckpoint(checkpoint)
			Expect(externalCheckpoint.Iterations).To(Equal(int32(10)))
			Expect(externalCheckpoint.Successes).To(Equal(int32(5)))
		})
	})
})
