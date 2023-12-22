package mostwater

func maxArea(height []int) int {
	l := len(height)

	if l == 0 {
		return 0
	}
	result := 0
	left, right := 0, len(height)-1

	for left < right {
		result = max(result, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(firsN, secondN int) int {
	if firsN > secondN {
		return secondN
	}
	return firsN
}

func max(firsN, secondN int) int {
	if firsN > secondN {
		return firsN
	}
	return secondN
}
