package mergedinterval

func merge(intervals [][]int) [][]int {

	for {
		result := processMerge(intervals)
		if len(result) == len(intervals) {
			return result
		}
		intervals = result
	}
}

func processMerge(intervals [][]int) [][]int {
	l := len(intervals)
	if l < 2 {
		return intervals
	}
	markedItem := make(map[int]bool, len(intervals))
	result := make([][]int, 0, len(intervals))
	index := 0
	for len(markedItem) < l {
		if ok := markedItem[index]; ok {
			index = index + 1
			continue
		}
		markedItem[index] = true
		start, end := intervals[index][0], intervals[index][1]
		for i := index + 1; i < len(intervals); i++ {

			if intervals[i][0] >= start && intervals[i][0] <= end {
				if intervals[i][1] >= end {
					end = intervals[i][1]
				}
				markedItem[i] = true
			} else if intervals[i][1] <= end && intervals[i][1] >= start {
				if intervals[i][0] <= start {
					start = intervals[i][0]
				}
				markedItem[i] = true
			} else if intervals[i][0] <= start && intervals[i][1] >= end {
				start = intervals[i][0]
				end = intervals[i][1]
				markedItem[i] = true
			}
		}
		result = append(result, []int{start, end})
		index = 0
	}
	return result
}
