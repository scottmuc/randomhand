package randomhand_test

import (
	"github.com/scottmuc/randomhand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Deal", func() {
	var deck randomhand.Deck

	BeforeEach(func() {
		deck = randomhand.Deck{
			randomhand.Cards["Ad"],
			randomhand.Cards["Ac"],
			randomhand.Cards["Ah"],
			randomhand.Cards["As"],
			randomhand.Cards["2d"],
		}
	})

	It("deals 5 cards from the top of the deck", func() {
		playerHand, _ := randomhand.Deal(deck)
		Expect(playerHand).To(HaveLen(5))
		Expect(playerHand).To(Equal(testHand("Ad", "Ac", "Ah", "As", "2d")))
	})

	It("removes the dealt cards from the deck", func() {
		_, newDeck := randomhand.Deal(deck)
		Expect(newDeck).To(HaveLen(0))
	})
})

// This test doesn't actually validate anything.
var _ = Describe("Shuffle", func() {
	It("changes the order of the cards", func() {
		deck := randomhand.NewStandardDeck()
		shuffledDeck := randomhand.Shuffle(deck)
		Expect(shuffledDeck).NotTo(Equal(randomhand.NewStandardDeck()))
	})
})
