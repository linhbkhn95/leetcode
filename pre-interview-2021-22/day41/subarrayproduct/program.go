package subarrayproduct

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0
	l := len(nums)
	total := 1
	for i := 0; i < l-1; i++ {
		total = nums[i]
		if total >= k {
			continue
		}
		result += 1
		for j := i + 1; j < l; j++ {
			total *= nums[j]
			if total < k {
				result++
			} else {
				break
			}
		}
	}
	if nums[l-1] < k {
		result += 1
	}
	return result
}
