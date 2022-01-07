package program

func findPeakElement(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	start := 0
	end := l - 1

	for start > -1 && start <= end {
		middle := (start + end) / 2
		if middle > 0 && middle < l-1 && middle < l && nums[middle] > nums[middle-1] && nums[middle] > nums[middle+1] {
			return middle
		} else if middle < l-1 && nums[middle] < nums[middle+1] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return start
}
