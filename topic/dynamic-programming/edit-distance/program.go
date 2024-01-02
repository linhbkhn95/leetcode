package editdistance

type key struct {
	m int
	n int
}

func minDistance(word1 string, word2 string) int {
	footprint := make(map[key]int)
	return dp(word1, word2, len(word1), len(word2), footprint)
}

func dp(word1, word2 string, m, n int, footprint map[key]int) int {
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	key := key{
		m: m,
		n: n,
	}
	if total, ok := footprint[key]; ok {
		return total
	}
	if word1[m-1] == word2[n-1] { // last word1 and word2 is the same
		return dp(word1, word2, m-1, n-1, footprint)
	}
	numOfOperations := 1 + min(
		dp(word1, word2, m, n-1, footprint),   // insert the last word1
		dp(word1, word2, m-1, n, footprint),   // remove the last word1
		dp(word1, word2, m-1, n-1, footprint), // replace the last word1
	)
	footprint[key] = numOfOperations
	return numOfOperations
}

func min(nums ...int) int {
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
