package integration_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"manabase-simulation/api"
	"manabase-simulation/package/validation"
)

const (
	validDeckList = `
Deck
5 Mountain
4 Manifold Mouse
4 Heartfire Hero
2 Screaming Nemesis
4 Emberheart Challenger
4 Monstrous Rage
2 Rockface Village
2 Obliterating Bolt
4 Thornspire Verge
2 Snakeskin Veil
3 Monastery Swiftspear
4 Burst Lightning
3 Soulstone Sanctuary
3 Questing Druid // Seek the Beast
4 Copperline Gorge
4 Hired Claw
2 Scorching Dragonfire
4 Karplusan Forest

Sideboard
3 Pawpatch Formation
2 Torch the Tower
2 Scorching Shot
2 Tectonic Hazard
1 Questing Druid
2 Lithomantic Barrage
3 Urabrask's Forge

`

	sampleDecklistWithInvalidCard = `
Deck
5 Mountain
3 Soulstone Sanctuary
3 Questing Druid

Sideboard
3 Pawpatch Formation

`
)

var _ = Describe("ValidateDeckList Integration Test", Label(IntegrationTestLabel), func() {

	It("Validates a valid deck list", func() {
		resp, err := Client.ValidateDeckList(context.Background(), &api.ValidateDeckListRequest{
			DeckList: validDeckList,
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.IsValid).To(BeTrue())
		Expect(resp.InvalidCards).To(HaveLen(0))
	})

	It("Validates an invalid deck list", func() {
		resp, err := Client.ValidateDeckList(context.Background(), &api.ValidateDeckListRequest{
			DeckList: sampleDecklistWithInvalidCard,
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.IsValid).To(BeFalse())
		Expect(resp.InvalidCards).To(HaveLen(1))
		Expect(resp.InvalidCards[0].Name).To(Equal("Questing Druid"))
		Expect(resp.InvalidCards[0].Reason).To(Equal(validation.MissingCardReason))
	})
})
