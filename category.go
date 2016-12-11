package randomhand

type Category struct {
	Name string
}

func Categorize(hand Hand) Category {
	if isRoyalFlush(hand) {
		return Category{Name: "RoyalFlush"}
	}
	if isFlush(hand) && isStraight(hand) {
		return Category{Name: "StraightFlush"}
	}
	if hasFourOfAKind(hand) {
		return Category{Name: "FourOfAKind"}
	}
	return Category{Name: "FullHouse"}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func hasFourOfAKind(hand Hand) bool {
	counter := map[string]int{}

	for _, card := range hand {
		counter[card.Rank] = counter[card.Rank] + 1
	}

	for _, count := range counter {
		if count == 4 {
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

	ranks := []string{"T", "J", "Q", "K", "A"}
	for _, card := range hand {
		if !contains(ranks, card.Rank) {
			return false
		}
	}

	return true
}
