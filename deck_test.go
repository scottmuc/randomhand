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
	var deck randomhand.Deck

	BeforeEach(func() {
		deck = randomhand.Deck{
			randomhand.Card{Rank: "A", Suit: "d"},
			randomhand.Card{Rank: "A", Suit: "c"},
			randomhand.Card{Rank: "A", Suit: "h"},
			randomhand.Card{Rank: "A", Suit: "s"},
			randomhand.Card{Rank: "2", Suit: "d"},
		}
	})

	It("deals 5 cards from the top of the deck", func() {
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
		_, newDeck := randomhand.Deal(deck)
		Expect(newDeck).To(HaveLen(0))
	})
})

// This is a horrible test, and it assumes that the
// Shuffle function is implementing some type of
// randomization
var _ = Describe("Shuffle", func() {
	It("changes the order of the cards", func() {
		deck := randomhand.Deck{
			randomhand.Card{Rank: "A", Suit: "d"},
			randomhand.Card{Rank: "A", Suit: "c"},
			randomhand.Card{Rank: "A", Suit: "h"},
			randomhand.Card{Rank: "A", Suit: "s"},
			randomhand.Card{Rank: "2", Suit: "d"},
		}

		shuffledDeck := randomhand.Shuffle(deck)

		Expect(shuffledDeck).NotTo(Equal(randomhand.Hand{
			randomhand.Card{Rank: "A", Suit: "d"},
			randomhand.Card{Rank: "A", Suit: "c"},
			randomhand.Card{Rank: "A", Suit: "h"},
			randomhand.Card{Rank: "A", Suit: "s"},
			randomhand.Card{Rank: "2", Suit: "d"},
		}))
	})
})

var _ = Describe("RoyalFlush", func() {
	It("ranks 10d,Jd,Qd,Kd,Ac as a royal flush", func() {
		hand := randomhand.Hand{
			randomhand.Card{Rank: "10", Suit: "d"},
			randomhand.Card{Rank: "J", Suit: "d"},
			randomhand.Card{Rank: "Q", Suit: "d"},
			randomhand.Card{Rank: "K", Suit: "d"},
			randomhand.Card{Rank: "A", Suit: "d"},
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("RoyalFlush"))
	})
})

var _ = Describe("StraightFlush", func() {
	It("ranks 7h,8h,9h,10h,Jh as a straight flush", func() {
		hand := randomhand.Hand{
			randomhand.Card{Rank: "7", Suit: "h"},
			randomhand.Card{Rank: "8", Suit: "h"},
			randomhand.Card{Rank: "9", Suit: "h"},
			randomhand.Card{Rank: "10", Suit: "h"},
			randomhand.Card{Rank: "J", Suit: "h"},
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("StraightFlush"))
	})
})
