package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	start, end := newInterval[0], newInterval[1]
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if intervals[len(intervals)-1][1] < start {
		return append(intervals, newInterval)
	}
	//1,3 2,4 2,4 1,3
	result := make([][]int, 0)
	isMerged := false
	for i, interval := range intervals {
		if start >= interval[0] && start <= interval[1] {
			start = interval[0]
			end = max(end, interval[1])
			intervals[i] = []int{start, end}
			if !isMerged {
				result = append(result, []int{start, end})
				isMerged = true
			} else {
				result[len(result)-1] = []int{start, end}
			}
		} else if interval[0] >= start && end >= interval[0] {
			end = max(end, interval[1])
			intervals[i] = []int{start, end}
			if !isMerged {
				result = append(result, []int{start, end})
				isMerged = true
			} else {
				result[len(result)-1] = []int{start, end}
			}
		} else if end < interval[0] {
			if !isMerged {
				result = append(result, []int{start, end})
			}
			result = append(result, intervals[i:]...)
			break
		} else {
			result = append(result, interval)
		}
	}

	return result
}
