package main

import (
	"fmt"

	"github.com/scottmuc/randomhand"
)

func main() {
	deck := randomhand.NewStandardDeck()

	playerHand, deck := randomhand.Deal(deck)
	computerHand, _ := randomhand.Deal(deck)

	fmt.Printf("Your hand: %v\n\n", playerHand)
	fmt.Printf("Computer's hand: %v\n\n", computerHand)
}
