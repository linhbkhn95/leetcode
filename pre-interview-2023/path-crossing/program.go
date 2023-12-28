package pathcrossing

import (
	"fmt"
)

func isPathCrossing(path string) bool {
	l := len(path)
	x, y := 0, 0
	footprint := make(map[string]struct{})
	footprint[fmt.Sprintf("%d_%d", x, y)] = struct{}{}
	for i := 0; i < l; i++ {
		switch path[i] {
		case 'N':
			{
				y++
			}
		case 'S':
			{
				y--
			}
		case 'W':
			{
				x--
			}
		case 'E':
			{
				x++
			}
		}
		fmt.Println(string(path[i]), x, y)

		if _, ok := footprint[fmt.Sprintf("%d_%d", x, y)]; ok {
			return true
		}
		footprint[fmt.Sprintf("%d_%d", x, y)] = struct{}{}
	}
	return false
}
