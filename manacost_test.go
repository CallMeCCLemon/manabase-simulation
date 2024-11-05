package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ManaCost", func() {
	When("Using a ManaCost", func() {
		It("Computes the remaining cost correctly when it has both generic and color-specific requirements", func() {
			cost := ManaCost{
				ColorRequirements: []ManaColor{white, white},
				GenericCost:       1,
			}
			Expect(cost.GetRemainingCost()).To(Equal(3))
		})

		It("Computes the remaining cost correctly when it only has generic requirements", func() {
			cost := ManaCost{
				ColorRequirements: []ManaColor{},
				GenericCost:       4,
			}
			Expect(cost.GetRemainingCost()).To(Equal(4))
		})

		It("Computes the remaining cost correctly when it only has color-specific requirements", func() {
			cost := ManaCost{
				ColorRequirements: []ManaColor{white, white},
				GenericCost:       0,
			}
			Expect(cost.GetRemainingCost()).To(Equal(2))
		})
	})
})
