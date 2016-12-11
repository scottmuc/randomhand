package randomhand

import (
	"math/rand"
)

type Card struct {
	Rank  string
	Suit  string
	Value int
}

type Deck []Card
type Hand []Card

var Cards = map[string]Card{
	// Diamonds
	"2d": Card{Rank: "2", Suit: "d", Value: 2},
	"3d": Card{Rank: "3", Suit: "d", Value: 3},
	"4d": Card{Rank: "4", Suit: "d", Value: 4},
	"5d": Card{Rank: "5", Suit: "d", Value: 5},
	"6d": Card{Rank: "6", Suit: "d", Value: 6},
	"7d": Card{Rank: "7", Suit: "d", Value: 7},
	"8d": Card{Rank: "8", Suit: "d", Value: 8},
	"9d": Card{Rank: "9", Suit: "d", Value: 9},
	"Td": Card{Rank: "T", Suit: "d", Value: 10},
	"Jd": Card{Rank: "J", Suit: "d", Value: 11},
	"Qd": Card{Rank: "Q", Suit: "d", Value: 12},
	"Kd": Card{Rank: "K", Suit: "d", Value: 13},
	"Ad": Card{Rank: "A", Suit: "d", Value: 14},
	// Clubs
	"2c": Card{Rank: "2", Suit: "c", Value: 2},
	"3c": Card{Rank: "3", Suit: "c", Value: 3},
	"4c": Card{Rank: "4", Suit: "c", Value: 4},
	"5c": Card{Rank: "5", Suit: "c", Value: 5},
	"6c": Card{Rank: "6", Suit: "c", Value: 6},
	"7c": Card{Rank: "7", Suit: "c", Value: 7},
	"8c": Card{Rank: "8", Suit: "c", Value: 8},
	"9c": Card{Rank: "9", Suit: "c", Value: 9},
	"Tc": Card{Rank: "T", Suit: "c", Value: 10},
	"Jc": Card{Rank: "J", Suit: "c", Value: 11},
	"Qc": Card{Rank: "Q", Suit: "c", Value: 12},
	"Kc": Card{Rank: "K", Suit: "c", Value: 13},
	"Ac": Card{Rank: "A", Suit: "c", Value: 14},
	// Hearts
	"2h": Card{Rank: "2", Suit: "h", Value: 2},
	"3h": Card{Rank: "3", Suit: "h", Value: 3},
	"4h": Card{Rank: "4", Suit: "h", Value: 4},
	"5h": Card{Rank: "5", Suit: "h", Value: 5},
	"6h": Card{Rank: "6", Suit: "h", Value: 6},
	"7h": Card{Rank: "7", Suit: "h", Value: 7},
	"8h": Card{Rank: "8", Suit: "h", Value: 8},
	"9h": Card{Rank: "9", Suit: "h", Value: 9},
	"Th": Card{Rank: "T", Suit: "h", Value: 10},
	"Jh": Card{Rank: "J", Suit: "h", Value: 11},
	"Qh": Card{Rank: "Q", Suit: "h", Value: 12},
	"Kh": Card{Rank: "K", Suit: "h", Value: 13},
	"Ah": Card{Rank: "A", Suit: "h", Value: 14},
	// Spades
	"2s": Card{Rank: "2", Suit: "s", Value: 2},
	"3s": Card{Rank: "3", Suit: "s", Value: 3},
	"4s": Card{Rank: "4", Suit: "s", Value: 4},
	"5s": Card{Rank: "5", Suit: "s", Value: 5},
	"6s": Card{Rank: "6", Suit: "s", Value: 6},
	"7s": Card{Rank: "7", Suit: "s", Value: 7},
	"8s": Card{Rank: "8", Suit: "s", Value: 8},
	"9s": Card{Rank: "9", Suit: "s", Value: 9},
	"Ts": Card{Rank: "T", Suit: "s", Value: 10},
	"Js": Card{Rank: "J", Suit: "s", Value: 11},
	"Qs": Card{Rank: "Q", Suit: "s", Value: 12},
	"Ks": Card{Rank: "K", Suit: "s", Value: 13},
	"As": Card{Rank: "A", Suit: "s", Value: 14},
}

// NewStandardDeck returns a standard 52 card deck
func NewStandardDeck() (deck Deck) {
	for _, value := range Cards {
		deck = append(deck, value)
	}
	return
}

func dealCard(deck Deck) (card Card, newDeck Deck) {
	card = deck[0]
	newDeck = deck[1:]
	return
}

// Deal deals a 5 card hand
func Deal(deck Deck) (hand Hand, newDeck Deck) {
	newDeck = deck
	var card Card
	for i := 0; i < 5; i++ {
		card, newDeck = dealCard(newDeck)
		hand = append(hand, card)
	}

	return
}

// Shuffle attempts to implement Fisher-Yates algorithm
// https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
func Shuffle(deck Deck) (shuffledDeck Deck) {
	shuffledDeck = deck
	length := len(deck)
	for i := 1; i < length; i++ {
		random := rand.Intn(i + 1)
		if i != random {
			shuffledDeck[random], shuffledDeck[i] = shuffledDeck[i], shuffledDeck[random]
		}
	}
	return
}
