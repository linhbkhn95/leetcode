package firstbadversion

func searchRange(nums []int, target int) []int {

	l := len(nums)

	start := 0
	end := l - 1

	for start <= end {
		middleOffset := (start + end) / 2
		if nums[middleOffset] == target {
			return extractFirtAndLastOffset(nums, middleOffset, target, l)
		} else if nums[middleOffset] < target {
			start = middleOffset + 1
		} else {
			end = middleOffset - 1
		}
	}
	return []int{-1, -1}
}

func extractFirtAndLastOffset(nums []int, currentOffset, target, l int) []int {
	i, j := currentOffset-1, currentOffset+1
	for i >= 0 && nums[i] == target {
		i--
	}
	for j < l && nums[j] == target {
		j++
	}
	return []int{i + 1, j - 1}
}
