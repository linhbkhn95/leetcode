package targetsum

func findTargetSumWays(nums []int, target int) int {
	footprint := make(map[key]int)
	return dp(nums, len(nums), 0, footprint, nums[0], target) + dp(nums, len(nums), 0, footprint, -nums[0], target)
}

type key struct {
	current int
	index   int
}

func dp(nums []int, l, index int, footprint map[key]int, current, target int) int {
	if index == l-1 {
		if current == target {
			return 1
		}
		return 0
	}
	k := key{
		current: current,
		index:   index,
	}
	if w, ok := footprint[k]; ok {
		return w
	}

	ways := dp(nums, l, index+1, footprint, current+nums[index+1], target) + dp(nums, l, index+1, footprint, current-nums[index+1], target)
	footprint[k] = ways
	return ways
}
