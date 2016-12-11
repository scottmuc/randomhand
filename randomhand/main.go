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

	fmt.Printf("Your hand (%s):\t\t%s\n", randomhand.Categorize(playerHand).Name, playerHand)
	fmt.Printf("Computer's hand (%s):\t%s\n", randomhand.Categorize(computerHand).Name, computerHand)
}
