package randomhand

type Card struct {
	Rank  string
	Value string
}

type Deck []Card

func NewStandardDeck() (deck Deck) {
	ranks := []string{"A", "2", "3", "4", "5", "6", "7",
		"8", "9", "T", "J", "Q", "K"}

	suits := []string{"d", "c", "h", "s"}

	for i := 0; i < len(ranks); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Rank:  suits[n],
				Value: ranks[i],
			}
			deck = append(deck, card)
		}
	}

	return
}
