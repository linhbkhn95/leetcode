package firstmissingpositive

import "math"

func firstMissingPositive(nums []int) int {
	minPositive := math.MaxInt32
	for _, n := range nums {
		if n > 0 {
			minPositive = min(minPositive, n)
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
