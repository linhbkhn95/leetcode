package revealcardsinincreasingorder

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	if len(deck) < 2 {
		return deck
	}
	sort.Ints(deck)
	result := make([]int, 0, len(deck))
	result = append(result, deck[len(deck)-2])
	result = append(result, deck[len(deck)-1])

	for i := len(deck) - 3; i >= 0; i-- {
		tail := result[len(result)-1]
		result = append([]int{deck[i], tail}, result[:len(result)-1]...)
	}
	return result
}
