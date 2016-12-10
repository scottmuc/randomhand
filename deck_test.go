package randomhand_test

import (
	"github.com/scottmuc/randomhand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Card", func() {
	It("has the standard 52 cards", func() {
		deck := randomhand.NewStandardDeck()

		ranks := []string{"A", "2", "3", "4", "5", "6", "7",
			"8", "9", "T", "J", "Q", "K"}

		suits := []string{"d", "c", "h", "s"}

		for i := 0; i < len(ranks); i++ {
			for n := 0; n < len(suits); n++ {
				sampleCard := randomhand.Card{
					Rank:  suits[n],
					Value: ranks[i],
				}
				Expect(deck).To(ContainElement(sampleCard))
			}
		}
	})
})
