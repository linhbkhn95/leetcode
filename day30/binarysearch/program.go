package program

func search(nums []int, target int) int {
	l := len(nums)

	start := 0
	end := l - 1

	for start <= end {
		middleOffset := (start + end) / 2
		if nums[middleOffset] == target {
			return middleOffset
		}
		if nums[middleOffset] > target {
			end = middleOffset - 1
		} else {
			start = middleOffset + 1
		}
	}
	return -1
}
