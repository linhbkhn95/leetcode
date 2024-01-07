package longestincreasingsubsequence

type key struct {
	index      int
	currentLen int
	maxLen     int
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func lengthOfLIS1(nums []int) int {
	footprint := make(map[int]key)
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		maxLen = max(maxLen, dp(nums, len(nums), i, nums[i], 1, footprint))
		footprint[i] = key{
			index:  i,
			maxLen: maxLen,
		}

	}
	return maxLen
}

func dp(nums []int, l, index, lastNumber, currentLen int, footprint map[int]key) int {
	if index == l {
		return 1
	}
	// k := key{
	// 	index:      index,
	// 	currentLen: currentLen,
	// }
	if v, ok := footprint[index]; ok {
		return v.maxLen
	}
	var maxLen int
	if nums[index] > lastNumber {
		maxLen = max(1+dp(nums, l, index+1, nums[index], currentLen+1, footprint), dp(nums, l, index+1, lastNumber, currentLen, footprint))
	} else {
		maxLen = dp(nums, l, index+1, lastNumber, currentLen, footprint)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
