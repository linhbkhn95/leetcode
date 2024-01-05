package widestverticalareabetweentwopointscontainingnopoints

import (
	"math"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var max int
	for i := 0; i < len(points)-1; i++ {
		verticalArea := int(math.Abs(float64(points[i][0] - points[i+1][0])))
		if verticalArea != 0 && verticalArea > max {
			max = verticalArea
		}
	}
	return max
}
