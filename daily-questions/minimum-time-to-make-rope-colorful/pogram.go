package minimumtimetomakeropecolorful

func minCost(colors string, neededTime []int) int {
	l := len(neededTime)
	lastIndex := 0
	sum, max := 0, neededTime[lastIndex]
	for i := 1; i < l; i++ {
		if colors[lastIndex] == colors[i] {
			if neededTime[i] > max {
				sum += max
				max = neededTime[i]
			} else {
				sum += neededTime[i]
			}
		} else {
			lastIndex = i
			max = neededTime[i]
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
