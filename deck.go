package randomhand

import (
	"math/rand"
)

type Card struct {
	Rank string
	Suit string
}

type Category struct {
	Name string
}

type Deck []Card
type Hand []Card

// NewStandardDeck returns a standard 52 card deck
func NewStandardDeck() (deck Deck) {
	ranks := []string{"A", "2", "3", "4", "5", "6", "7",
		"8", "9", "T", "J", "Q", "K"}

	suits := []string{"d", "c", "h", "s"}

	for i := 0; i < len(ranks); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Suit: suits[n],
				Rank: ranks[i],
			}
			deck = append(deck, card)
		}
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

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isFlush(hand Hand) bool {
	suit := hand[0].Suit
	for _, card := range hand {
		if suit != card.Suit {
			return false
		}
	}
	return true
}

func isStraight(hand Hand) bool {
	return true
}

func isRoyalFlush(hand Hand) bool {
	if !isFlush(hand) {
		return false
	}

	ranks := []string{"10", "J", "Q", "K", "A"}
	for _, card := range hand {
		if !contains(ranks, card.Rank) {
			return false
		}
	}

	return true
}

func Categorize(hand Hand) Category {
	if isRoyalFlush(hand) {
		return Category{Name: "RoyalFlush"}
	}
	if isFlush(hand) && isStraight(hand) {
		return Category{Name: "StraightFlush"}
	}
	return Category{Name: "FourOfAKind"}
}
