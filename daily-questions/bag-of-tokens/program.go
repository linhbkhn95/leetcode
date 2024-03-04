package bagoftokens

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	start, end := 0, len(tokens)-1
	result := 0
	debit := 0
	for start <= end {
		if power >= tokens[start] {
			result++
			power -= tokens[start]
			start++
			debit = 0
		} else if result > 0 {
			power += tokens[end]
			result--
			end--
			debit++
		} else {
			break
		}

	}
	if debit > 0 {
		return result + debit
	}
	return result
}
