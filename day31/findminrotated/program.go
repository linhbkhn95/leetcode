package program

func findMin(nums []int) int {

	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return min(nums[0], nums[1])
	}
	start := 0
	end := l - 1

	for start <= end {
		if (end - start) == 1 {
			return min(nums[start], nums[end])
		}
		middle := (end + start) / 2

		if middle >= 1 && middle < l-1 && nums[middle-1] > nums[middle] && nums[middle] < nums[middle+1] {
			return nums[middle]
		}

		if nums[middle] >= nums[start] && nums[middle] <= nums[end] {
			return nums[start]
		} else if nums[middle] > nums[start] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}

func min(first, second int) int {
	if first > second {
		return second
	}
	return first
}

func findMin1(nums []int) int {

	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return min(nums[0], nums[1])
	}
	start := 0
	end := l - 1

	for start < end {
		if (end - start) == 1 {
			return min(nums[start], nums[end])
		}
		middle := (end + start) / 2

		if nums[middle] > nums[end] {
			start = middle + 1
		} else {
			end = middle
		}

	}
	return nums[start]
}
