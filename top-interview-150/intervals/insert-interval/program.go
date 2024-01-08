package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if intervals[0][0] > newInterval[1] {
		return append([][]int{newInterval}, intervals...)
	}
	if intervals[len(intervals)-1][1] < newInterval[0] {
		return append(intervals, newInterval)
	}
	result := make([][]int, 0)
	start, end := newInterval[0], newInterval[1]
	isMerged := false
	for i, interval := range intervals {
		if start > interval[1] || end < interval[0] {
			if !isMerged && i > 0 && intervals[i-1][1] < start && interval[0] > end {
				result = append(result, newInterval)
				result = append(result, intervals[i:]...)
				break
			}
			result = append(result, interval)
		} else {
			if interval[0] <= start && interval[1] >= start {
				start = interval[0]
			}
			end = max(end, interval[1])
			if !isMerged {
				result = append(result, []int{start, end})
				isMerged = true
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
