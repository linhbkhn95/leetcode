package minimumburstballons

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var result int

	i := 1
	raceCondition := points[0]
	result = 1
	for i < len(points) {
		if points[i][0] <= raceCondition[1] {
			raceCondition[0] = points[i][0]
			if points[i][1] <= raceCondition[1] {
				raceCondition[1] = points[i][1]
			}
		} else {
			result++
			raceCondition = points[i]
		}
		i++
	}
	return result
}
