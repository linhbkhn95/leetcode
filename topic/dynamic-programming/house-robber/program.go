package houserobber

func rob1(nums []int) int {
	l := len(nums)
	footprint := make(map[key]int)
	return max(dp(nums, l, 0, true, footprint), dp(nums, l, 0, false, footprint))
}

func rob(nums []int) int {
	l := len(nums)
	footprint := make(map[int]int)
	return doRob(nums, l, 0, 0, 0, footprint)
}

func doRob(nums []int, l, index, currentSum, largestSum int, footprint map[int]int) int {
	if index > l-1 {
		return max(currentSum, largestSum)
	}
	value, ok := footprint[index]
	if ok && value >= currentSum {
		return largestSum
	}
	footprint[index] = currentSum
	return max(doRob(nums, l, index+2, currentSum+nums[index], largestSum, footprint), doRob(nums, l, index+1, currentSum, largestSum, footprint))
}

type key struct {
	index int
	isRob bool
}

func dp(nums []int, l, index int, isRob bool, footprint map[key]int) int {
	if index > l-1 {
		return 0
	}
	key := key{
		index: index,
		isRob: isRob,
	}
	value, ok := footprint[key]
	if ok {
		return value
	}
	var maxV int
	if isRob {
		maxV = max(nums[index]+dp(nums, l, index+2, true, footprint), nums[index]+dp(nums, l, index+2, false, footprint))
	} else {
		maxV = max(dp(nums, l, index+1, true, footprint), dp(nums, l, index+1, false, footprint))
	}
	footprint[key] = maxV
	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
