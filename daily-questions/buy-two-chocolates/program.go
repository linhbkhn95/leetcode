package buytwochocolates

import "fmt"

func buyChoco(prices []int, money int) int {
	min1, min2 := prices[0], prices[1]
	for i := 2; i < len(prices); i++ {
		if prices[i] < min1 {
			if min1 < min2 {
				min2 = prices[i]
				continue
			}
			min1 = prices[i]
		} else if prices[i] < min2 {
			if min2 < min1 {
				min1 = prices[i]
				continue
			}
			min2 = prices[i]
		}
	}
	fmt.Println("min1,min2", min1, min2)
	leftover := money - (min1 + min2)
	if leftover < 0 {
		return money
	}
	return leftover
}
