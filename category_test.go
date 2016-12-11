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
