package randomhand_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/scottmuc/randomhand"
)

var _ = Describe("RoyalFlush", func() {
	It("ranks Td,Jd,Qd,Kd,Ad as a royal flush", func() {
		hand := testHand("Td", "Jd", "Qd", "Kd", "Ad")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("RoyalFlush"))
	})
})

var _ = Describe("StraightFlush", func() {
	It("ranks 7h,8h,9h,Th,Jh as a straight flush", func() {
		hand := testHand("7h", "8h", "9h", "Th", "Jh")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("StraightFlush"))
	})
})

var _ = Describe("FourOfAKind", func() {
	It("ranks 5h,5d,5c,5s,Kh as a four of a kind", func() {
		hand := testHand("5h", "5d", "5c", "5s", "Kh")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("FourOfAKind"))
	})
})

var _ = Describe("FullHouse", func() {
	It("ranks 9h,9d,9c,8s,8h as a four of a kind", func() {
		hand := testHand("9h", "9d", "9c", "8s", "8h")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("FullHouse"))
	})
})

var _ = Describe("Flush", func() {
	It("ranks 3h,8h,Qh,2h,9h as a flush", func() {
		hand := testHand("3h", "8h", "Qh", "2h", "9h")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("Flush"))
	})
})

var _ = Describe("Straight", func() {
	It("ranks 3h,4c,5h,6h,7s as a straight", func() {
		hand := testHand("3h", "4c", "5h", "6h", "7s")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("Straight"))
	})
})

var _ = Describe("ThreeOfAKind", func() {
	It("ranks 3h,3c,3s,6h,7s as a three of a kind", func() {
		hand := testHand("3h", "3c", "3s", "6h", "7s")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("ThreeOfAKind"))
	})
})

var _ = Describe("TwoPair", func() {
	It("ranks 3h,3c,6s,6h,7s as two pair", func() {
		hand := testHand("3h", "3c", "6s", "6h", "7s")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("TwoPair"))
	})
})

var _ = Describe("OnePair", func() {
	It("ranks 3h,3c,9s,6h,7s as one pair", func() {
		hand := testHand("3h", "3c", "9s", "6h", "7s")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("OnePair"))
	})
})

var _ = Describe("HighCard", func() {
	It("ranks 3h,2c,9s,6h,7s as high card", func() {
		hand := testHand("3h", "2c", "9s", "6h", "7s")
		category := randomhand.Categorize(hand)
		Expect(category.Name).To(Equal("HighCard"))
	})
})
