package main

import "math"

func minSumOfLengths(arr []int, target int) int {
	l := len(arr)
	dp := [10001]int{}
	start, end := 0, 0
	sum := 0
	minLen := math.MaxInt32
	result := math.MaxInt32
	for end < l {
		sum += arr[end]
		for sum > target {
			sum -= arr[start]
			start++
		}
		if sum == target {
			len := end - start + 1
			minLen = min(minLen, len)
			if start > 0 {
				result = min(dp[start-1]+len, result)
			}
		}
		dp[end] = minLen
		end++
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}
