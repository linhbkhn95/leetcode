package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	start, end := newInterval[0], newInterval[1]
	isMerged := false
	for _, interval := range intervals {
		if start > interval[1] || end < interval[1] {
			result = append(result, interval)
		} else {
			if !isMerged && start <= interval[0] {
			}
			if interval[0] <= start && interval[1] >= start {
				start = interval[0]
			}
			end = max(end, interval[1])
			if !isMerged {
				result = append(result, []int{start, end})
			} else {
				result[len(result)-1] = []int{start, end}
			}
		}
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
