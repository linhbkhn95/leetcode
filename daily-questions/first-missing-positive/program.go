package first_missing_positive

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num <= 0 || num > n {
			nums[i] = n + 1
		}
	}
	for _, num := range nums {
		if num < 0 || num >= n || nums[num] == -1 {
			continue
		}
		nums[num-1] *= -1
	}
	for i, num := range nums {
		if num >= n {
			return i + 1
		}
	}
	return n
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
