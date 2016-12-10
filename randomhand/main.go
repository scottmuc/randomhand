package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/scottmuc/randomhand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	deck := randomhand.NewStandardDeck()
	shuffledDeck := randomhand.Shuffle(deck)

	playerHand, shuffledDeck := randomhand.Deal(shuffledDeck)
	computerHand, _ := randomhand.Deal(shuffledDeck)

	fmt.Printf("Your hand: %v\n\n", playerHand)
	fmt.Printf("Computer's hand: %v\n\n", computerHand)
}
