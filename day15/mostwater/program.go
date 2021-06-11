package mostwater

func maxArea(height []int) int {
	l := len(height)

	if l == 0 {
		return 0
	}
	result := 0
	for i := 0; i < l-1; i++ {
		if height[i+1] < height[i] {
			continue
		}
		for j := i + 1; j < l; j++ {
			area := (j - i) * min(height[i], height[j])
			if result < area {
				result = area
			}
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
