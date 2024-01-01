package houserobberii

type key struct {
	index            int
	includeZeroIndex bool
	isRobbed         bool
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l := len(nums)
	footprint := make(map[key]int)
	return max(dp(nums, l, true, true, 0, footprint), dp(nums, l, false, false, 0, footprint))
}

func dp(nums []int, l int, includeZeroIndex, isRobbed bool, index int, footprint map[key]int) int {
	if index > l-1 {
		footprint[getKey(index, isRobbed, includeZeroIndex)] = 0
		return 0
	}
	value, ok := footprint[getKey(index, isRobbed, includeZeroIndex)]
	if ok {
		return value
	}
	var maxV int
	if isRobbed {
		if index == l-1 {
			if includeZeroIndex {
				return 0
			}
			return nums[l-1]
		}
		maxV = max(nums[index]+dp(nums, l, includeZeroIndex, true, index+2, footprint), nums[index]+dp(nums, l, includeZeroIndex, false, index+2, footprint))
	} else {
		maxV = max(dp(nums, l, includeZeroIndex, true, index+1, footprint), dp(nums, l, includeZeroIndex, false, index+1, footprint))
	}
	footprint[getKey(index, isRobbed, includeZeroIndex)] = maxV
	return maxV
}

func getKey(index int, isRobbed, includeZeroIndex bool) key {
	return key{
		index:            index,
		includeZeroIndex: includeZeroIndex,
		isRobbed:         isRobbed,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
