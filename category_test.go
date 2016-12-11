package randomhand_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/scottmuc/randomhand"
)

var _ = Describe("RoyalFlush", func() {
	It("ranks Td,Jd,Qd,Kd,Ac as a royal flush", func() {
		hand := randomhand.Hand{
			randomhand.Cards["Td"],
			randomhand.Cards["Jd"],
			randomhand.Cards["Qd"],
			randomhand.Cards["Kd"],
			randomhand.Cards["Ad"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("RoyalFlush"))
	})
})

var _ = Describe("StraightFlush", func() {
	It("ranks 7h,8h,9h,10h,Jh as a straight flush", func() {
		hand := randomhand.Hand{
			randomhand.Cards["7h"],
			randomhand.Cards["8h"],
			randomhand.Cards["9h"],
			randomhand.Cards["Th"],
			randomhand.Cards["Jh"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("StraightFlush"))
	})
})

var _ = Describe("FourOfAKind", func() {
	It("ranks 5h,5d,5c,5s,Kh as a four of a kind", func() {
		hand := randomhand.Hand{
			randomhand.Cards["5d"],
			randomhand.Cards["5c"],
			randomhand.Cards["5h"],
			randomhand.Cards["5s"],
			randomhand.Cards["Kh"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("FourOfAKind"))
	})
})

var _ = Describe("FullHouse", func() {
	It("ranks 9h,9d,9c,8s,8h as a four of a kind", func() {
		hand := randomhand.Hand{
			randomhand.Cards["9h"],
			randomhand.Cards["9d"],
			randomhand.Cards["9c"],
			randomhand.Cards["8s"],
			randomhand.Cards["8h"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("FullHouse"))
	})
})

var _ = Describe("Flush", func() {
	It("ranks 3h,8h,Qh,2h,9h as a flush", func() {
		hand := randomhand.Hand{
			randomhand.Cards["3h"],
			randomhand.Cards["8h"],
			randomhand.Cards["Qh"],
			randomhand.Cards["2h"],
			randomhand.Cards["9h"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("Flush"))
	})
})

var _ = Describe("Straight", func() {
	It("ranks 3h,4c,5h,6h,7s as a straight", func() {
		hand := randomhand.Hand{
			randomhand.Cards["3h"],
			randomhand.Cards["4c"],
			randomhand.Cards["5h"],
			randomhand.Cards["6h"],
			randomhand.Cards["7s"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("Straight"))
	})
})

var _ = Describe("ThreeOfAKind", func() {
	It("ranks 3h,3c,3s,6h,7s as a three of a kind", func() {
		hand := randomhand.Hand{
			randomhand.Cards["3h"],
			randomhand.Cards["3c"],
			randomhand.Cards["3s"],
			randomhand.Cards["6h"],
			randomhand.Cards["7s"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("ThreeOfAKind"))
	})
})

var _ = Describe("TwoPair", func() {
	It("ranks 3h,3c,6s,6h,7s as two pair", func() {
		hand := randomhand.Hand{
			randomhand.Cards["3h"],
			randomhand.Cards["3c"],
			randomhand.Cards["6s"],
			randomhand.Cards["6h"],
			randomhand.Cards["7s"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("TwoPair"))
	})
})

var _ = Describe("OnePair", func() {
	It("ranks 3h,3c,9s,6h,7s as two pair", func() {
		hand := randomhand.Hand{
			randomhand.Cards["3h"],
			randomhand.Cards["3c"],
			randomhand.Cards["9s"],
			randomhand.Cards["6h"],
			randomhand.Cards["7s"],
		}
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("OnePair"))
	})
})
