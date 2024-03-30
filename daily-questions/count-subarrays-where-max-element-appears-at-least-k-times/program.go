package count_subarrays_where_max_element_appears_at_least_k_times

func countSubarrays(nums []int, k int) int64 {
	l := len(nums)
	maxN := getMax(nums)
	maxPositions := []int{}
	for i, n := range nums {
		if n == maxN {
			maxPositions = append(maxPositions, i)
		}
	}
	var result int64
	if len(maxPositions) < k {
		return 0
	}
	start, end := 0, 0
	end = start + k - 1
	offset := maxPositions[end]
	result += int64((maxPositions[start] + 1) * (l - offset))
	for start = 1; start < len(maxPositions)-k+1; start++ {
		end = start + k - 1
		offset = maxPositions[end]
		result += int64((maxPositions[start] - maxPositions[start-1]) * (l - offset))
	}
	return result
}

func getMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxN := nums[0]
	for i := 0; i < len(nums); i++ {
		maxN = max(maxN, nums[i])
	}
	return maxN
}

// 1,3,2,3,3           k 2
// 2 + 2 + 1 +1
// 1,3: 1-0+ 5-3 = 3, 2-4: 2-1,
//
