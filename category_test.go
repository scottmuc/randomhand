package randomhand_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/scottmuc/randomhand"
)

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

var _ = Describe("FourOfAKind", func() {
	It("ranks 5h,5d,5c,5s,Kh as a four of a kind", func() {
		hand := randomhand.Hand{
			randomhand.Card{Rank: "5", Suit: "d"},
			randomhand.Card{Rank: "5", Suit: "c"},
			randomhand.Card{Rank: "5", Suit: "h"},
			randomhand.Card{Rank: "5", Suit: "s"},
			randomhand.Card{Rank: "K", Suit: "h"},
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("FourOfAKind"))
	})
})
