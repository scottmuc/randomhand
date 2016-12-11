package randomhand

import (
	"math/rand"
)

type Card struct {
	Rank string
	Suit string
}

type Deck []Card
type Hand []Card

var Cards = map[string]Card{
	// Diamonds
	"2d": Card{Rank: "2", Suit: "d"},
	"3d": Card{Rank: "3", Suit: "d"},
	"4d": Card{Rank: "4", Suit: "d"},
	"5d": Card{Rank: "5", Suit: "d"},
	"6d": Card{Rank: "6", Suit: "d"},
	"7d": Card{Rank: "7", Suit: "d"},
	"8d": Card{Rank: "8", Suit: "d"},
	"9d": Card{Rank: "9", Suit: "d"},
	"Td": Card{Rank: "T", Suit: "d"},
	"Jd": Card{Rank: "J", Suit: "d"},
	"Qd": Card{Rank: "Q", Suit: "d"},
	"Kd": Card{Rank: "K", Suit: "d"},
	"Ad": Card{Rank: "A", Suit: "d"},
	// Clubs
	"2c": Card{Rank: "2", Suit: "c"},
	"3c": Card{Rank: "3", Suit: "c"},
	"4c": Card{Rank: "4", Suit: "c"},
	"5c": Card{Rank: "5", Suit: "c"},
	"6c": Card{Rank: "6", Suit: "c"},
	"7c": Card{Rank: "7", Suit: "c"},
	"8c": Card{Rank: "8", Suit: "c"},
	"9c": Card{Rank: "9", Suit: "c"},
	"Tc": Card{Rank: "T", Suit: "c"},
	"Jc": Card{Rank: "J", Suit: "c"},
	"Qc": Card{Rank: "Q", Suit: "c"},
	"Kc": Card{Rank: "K", Suit: "c"},
	"Ac": Card{Rank: "A", Suit: "c"},
	// Hearts
	"2h": Card{Rank: "2", Suit: "h"},
	"3h": Card{Rank: "3", Suit: "h"},
	"4h": Card{Rank: "4", Suit: "h"},
	"5h": Card{Rank: "5", Suit: "h"},
	"6h": Card{Rank: "6", Suit: "h"},
	"7h": Card{Rank: "7", Suit: "h"},
	"8h": Card{Rank: "8", Suit: "h"},
	"9h": Card{Rank: "9", Suit: "h"},
	"Th": Card{Rank: "T", Suit: "h"},
	"Jh": Card{Rank: "J", Suit: "h"},
	"Qh": Card{Rank: "Q", Suit: "h"},
	"Kh": Card{Rank: "K", Suit: "h"},
	"Ah": Card{Rank: "A", Suit: "h"},
	// Spades
	"2s": Card{Rank: "2", Suit: "s"},
	"3s": Card{Rank: "3", Suit: "s"},
	"4s": Card{Rank: "4", Suit: "s"},
	"5s": Card{Rank: "5", Suit: "s"},
	"6s": Card{Rank: "6", Suit: "s"},
	"7s": Card{Rank: "7", Suit: "s"},
	"8s": Card{Rank: "8", Suit: "s"},
	"9s": Card{Rank: "9", Suit: "s"},
	"Ts": Card{Rank: "T", Suit: "s"},
	"Js": Card{Rank: "J", Suit: "s"},
	"Qs": Card{Rank: "Q", Suit: "s"},
	"Ks": Card{Rank: "K", Suit: "s"},
	"As": Card{Rank: "A", Suit: "s"},
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
