package minimum_number_of_arrows_to_burst_balloons

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	result := 1
	start, end := points[0][0], points[0][1]
	for i := 0; i < len(points); i++ {
		point := points[i]
		if point[0] > end {
			result++
			start = point[0]
			end = point[1]
		} else {
			start = max(start, point[0])
			end = min(end, point[1])
		}
	}
	return result
}

// [1,6] [2,8] [7,12] [10, 16]
//[2,6] [7, 8] [10,12]
