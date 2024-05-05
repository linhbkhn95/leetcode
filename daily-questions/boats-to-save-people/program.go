package boatstosavepeople

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})
	result := 1
	currentSum := people[0]

	left, right := 0, len(people)-1
	count := 1
	for left < right {
		if currentSum+people[right] <= limit && count < 2 {
			currentSum += people[right]
			count++
			right--
		} else {
			result++
			left++
			currentSum = people[left]
			count = 1
		}
	}

	return result
}

// 1 2 4 5
// 50, (49), (47,2), (41,2,2), (33,6,6), (26,7,10), (22,10,10),(18,11,12),(13)
