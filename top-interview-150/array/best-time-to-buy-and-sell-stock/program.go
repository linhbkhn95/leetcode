package main

func maxProfit(prices []int) int {
	result := 0
	start := 0
	for start < len(prices)-1 {
		if prices[start+1]-prices[start] <= 0 {
			start++
			continue
		}
		for j := start + 1; j < len(prices); j++ {
			profit := prices[j] - prices[start]
			if profit > 0 && result < profit {
				result = profit
			}
		}
		i := start + 1
		for i < len(prices) && prices[i] >= prices[start] {
			i++
		}
		start = i
	}
	return result
}
