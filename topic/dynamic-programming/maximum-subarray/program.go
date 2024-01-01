package maximumsubarray

// Kadane algorithm to solve maximum subarray problem
// reference: https://en.wikipedia.org/wiki/Maximum_subarray_problem
func maxSubArray1(nums []int) int {
	largestSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if currentSum = max(nums[i], currentSum+nums[i]); currentSum > largestSum {
			largestSum = currentSum
		}
	}
	return largestSum
}

func maxSubArray(nums []int) int {
	dps := make([]int, len(nums))
	dps[0] = nums[0]
	maxSub := nums[0]
	for i := 1; i < len(nums); i++ {
		dps[i] = max(nums[i], dps[i-1]+nums[i])
		maxSub = max(maxSub, dps[i])
	}
	return maxSub
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
