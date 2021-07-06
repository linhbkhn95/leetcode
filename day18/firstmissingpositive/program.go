package firstmissingpositive

import "fmt"

func firstMissingPositive(nums []int) int {

	minPositive := 0

	positiveMapping := make(map[int]bool)
	for _, n := range nums {
		if n > 0 {
			positiveMapping[n] = true

			if minPositive > n || minPositive == 0 {
				minPositive = n
			}
		}
	}
	fmt.Println("minPositive", minPositive)
	if minPositive == 0 || minPositive > 1 {
		return 1
	}

	i := 1
	for {
		if _, ok := positiveMapping[minPositive+i]; !ok {
			return minPositive + i
		}
		i++
	}
}
