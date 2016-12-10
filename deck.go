package randomhand

type Card struct {
	Suit string
	Rank string
}

type Deck []Card
type Hand []Card

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

func Deal(deck Deck) (hand Hand, newDeck Deck) {
	newDeck = deck
	for i := 0; i < 5; i++ {
		hand = append(hand, newDeck[0])
		newDeck = newDeck[1:]
	}

	return
}
