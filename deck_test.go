package randomhand_test

import (
	"github.com/scottmuc/randomhand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deck", func() {
	It("has the standard 52 cards", func() {
		deck := randomhand.NewStandardDeck()

		ranks := []string{"A", "2", "3", "4", "5", "6", "7",
			"8", "9", "T", "J", "Q", "K"}

		suits := []string{"d", "c", "h", "s"}

		for i := 0; i < len(ranks); i++ {
			for n := 0; n < len(suits); n++ {
				sampleCard := randomhand.Card{
					Suit: suits[n],
					Rank: ranks[i],
				}
				Expect(deck).To(ContainElement(sampleCard))
			}
		}
	})
})

var _ = Describe("Deal", func() {
	It("deals 5 cards from the top of the deck", func() {
		deck := randomhand.Deck{
			randomhand.Card{Rank: "A", Suit: "d"},
			randomhand.Card{Rank: "A", Suit: "c"},
			randomhand.Card{Rank: "A", Suit: "h"},
			randomhand.Card{Rank: "A", Suit: "s"},
			randomhand.Card{Rank: "2", Suit: "d"},
		}

		playerHand, _ := randomhand.Deal(deck)

		Expect(playerHand).To(HaveLen(5))
		Expect(playerHand).To(Equal(randomhand.Hand{
			randomhand.Card{Rank: "A", Suit: "d"},
			randomhand.Card{Rank: "A", Suit: "c"},
			randomhand.Card{Rank: "A", Suit: "h"},
			randomhand.Card{Rank: "A", Suit: "s"},
			randomhand.Card{Rank: "2", Suit: "d"},
		}))
	})

	It("removes the dealt cards from the deck", func() {
		deck := randomhand.Deck{
			randomhand.Card{Rank: "A", Suit: "d"},
			randomhand.Card{Rank: "A", Suit: "c"},
			randomhand.Card{Rank: "A", Suit: "h"},
			randomhand.Card{Rank: "A", Suit: "s"},
			randomhand.Card{Rank: "2", Suit: "d"},
		}

		_, newDeck := randomhand.Deal(deck)

		Expect(newDeck).To(HaveLen(0))
	})
})
