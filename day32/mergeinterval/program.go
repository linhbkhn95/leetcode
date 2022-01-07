package program

import "sort"

func merge(intervals [][]int) [][]int {

	result := make([][]int, 0, len(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		l := len(result)
		k := l - 1
		for k >= 0 && result[k][0] >= intervals[i][0] && result[k][1] <= intervals[i][1] {
			k--
		}
		if k == -1 {
			result = [][]int{intervals[i]}
			continue
		}

		result = result[:k+1]
		var pre []int
		if k != -1 {
			pre = result[k]
		}
		if pre != nil && pre[1] >= intervals[i][0] {
			end := pre[1]
			if intervals[i][1] >= pre[1] {
				end = intervals[i][1]
			}
			result = result[:k]
			result = append(result, []int{pre[0], end})
		} else {
			result = append(result, intervals[i])
		}

	}
	return result
}
