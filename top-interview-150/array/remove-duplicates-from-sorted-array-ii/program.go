package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	result := 1
	pre := nums[0]
	i := 1
	nDuplicate := 1
	for i < len(nums) {
		if nums[i] == pre {
			nDuplicate += 1
			if nDuplicate == 2 {
				nums[result] = nums[i]
				result += 1
			}
			i++

		} else {
			pre = nums[i]
			nums[result] = nums[i]
			result += 1
			nDuplicate = 1
			i++
		}
	}
	return result
}
