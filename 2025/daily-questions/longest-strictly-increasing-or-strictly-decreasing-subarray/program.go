package main

func longestMonotonicSubarray(nums []int) int {
	result, countIncrease, countDecrease := 1, 1, 1
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < pre {
			result = max(result, countIncrease)
			countDecrease++
			countIncrease = 1
		} else if nums[i] > pre {
			result = max(result, countDecrease)
			countIncrease++
			countDecrease = 1
		} else {
			result = max(result, countIncrease)
			result = max(result, countDecrease)
			countDecrease, countIncrease = 1, 1
		}
		pre = nums[i]
	}
	result = max(result, countIncrease)
	result = max(result, countDecrease)
	return result
}
