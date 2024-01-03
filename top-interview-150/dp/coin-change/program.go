package coinchange

import (
	"sort"
)
import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c < 0 {
				continue
			}

			dp[a] = min(dp[a-c]+1, dp[a])
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 1
	}
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	footprint := make(map[key]int)
	result, ok := dp(coins, len(coins), 0, amount, footprint)
	if !ok {
		return -1
	}
	return result
}

type key struct {
	index  int
	amount int
}

func dp(coins []int, l, index, amount int, footprint map[key]int) (int, bool) {
	if l-1 == index {
		if amount < coins[index] || amount%coins[index] != 0 {
			return 0, false
		}

		return amount / coins[index], true
	}
	if amount < 0 {
		return 0, false
	}
	if amount == 0 {
		return 0, true
	}

	k := key{
		index:  index,
		amount: amount,
	}

	if v, ok := footprint[k]; ok {
		if v == 0 {
			return 0, false
		}
		return v, true
	}
	var count int
	count1, ok := dp(coins, l, index, amount-coins[index], footprint)
	if ok {
		count = count1 + 1
	}
	count2, ok := dp(coins, l, index+1, amount, footprint)
	if ok {
		if count == 0 {
			count = count2
		}
		count = min(count, count2)
	}
	footprint[k] = count
	if count == 0 {
		return 0, false
	}

	return count, true

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
