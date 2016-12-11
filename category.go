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
	if hasThreeOfAKind(hand) && hasAPair(hand) {
		return Category{Name: "FullHouse"}
	}
	if isFlush(hand) {
		return Category{Name: "Flush"}
	}
	if isStraight(hand) {
		return Category{Name: "Straight"}
	}
	if hasThreeOfAKind(hand) {
		return Category{Name: "ThreeOfAKind"}
	}
	return Category{Name: "TwoPair"}
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

func hasThreeOfAKind(hand Hand) bool {
	counter := map[string]int{}

	for _, card := range hand {
		counter[card.Rank] = counter[card.Rank] + 1
	}

	for _, count := range counter {
		if count == 3 {
			return true
		}
	}
	return false
}

func hasAPair(hand Hand) bool {
	counter := map[string]int{}

	for _, card := range hand {
		counter[card.Rank] = counter[card.Rank] + 1
	}

	for _, count := range counter {
		if count == 2 {
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

func containsInt(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isStraight(hand Hand) bool {
	values := []int{}
	for _, card := range hand {
		values = append(values, card.Value)
	}
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	for i := 0; i < len(values); i++ {
		if !containsInt(values, min+i) {
			return false
		}
	}

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
