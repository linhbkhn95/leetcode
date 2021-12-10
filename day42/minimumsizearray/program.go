package minimumsizearray

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	result := l + 1
	preAbility := true
	for i := range nums {
		if !preAbility {
			return 0
		}
		lR := backtracking(nums, l, i, 0, 0, result, target)
		if lR == l+1 {
			preAbility = false
		}
		result = min(result, lR)
	}
	if result == l+1 {
		return 0
	}
	return result
}

func backtracking(nums []int, l, start, currentLen, currentValue, minLen, targer int) int {
	if currentLen >= minLen-1 && minLen != l+1 {
		return minLen
	}
	if start == l-1 {
		if currentValue+nums[start] >= targer {
			return min(currentLen+1, minLen)
		}
		return minLen
	}
	if currentValue+nums[start] >= targer {
		return min(currentLen+1, minLen)
	}
	return backtracking(nums, l, start+1, currentLen+1, currentValue+nums[start], minLen, targer)
}

func min(first, second int) int {
	if first > second {
		return second
	}
	return first
}
