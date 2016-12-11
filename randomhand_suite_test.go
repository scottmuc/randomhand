package randomhand_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/scottmuc/randomhand"

	"testing"
)

func testHand(cards ...string) (hand randomhand.Hand) {
	for _, card := range cards {
		hand = append(hand, randomhand.Cards[card])
	}
	return
}

func TestRandomhand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Randomhand Suite")
}
