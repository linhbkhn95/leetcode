package consecutiveone

func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	l := len(nums)
	i := 0
	currentCount := 0
	for i < l {
		if nums[i] == 0 {
			if currentCount > result {
				result = currentCount
			}
			j := i + 1
			for j < l && nums[j] == 0 {
				j++
			}
			i = j
			currentCount = 0
		} else {
			currentCount++
			i++
		}
	}
	if currentCount > result {
		return currentCount
	}
	return result
}
