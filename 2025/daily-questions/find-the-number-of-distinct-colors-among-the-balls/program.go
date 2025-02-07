package main

func queryResults(limit int, queries [][]int) []int {
	balls := make(map[int]int)
	colors := make(map[int]int)
	result := make([]int, len(queries))
	for i, query := range queries {
		ball, newColor := query[0], query[1]
		oldColor, ok := balls[ball]
		if ok {
			colors[oldColor] -= 1
		}

		colors[newColor] += 1
		balls[ball] = newColor

		if colors[oldColor] == 0 {
			delete(colors, oldColor)
		}
		result[i] = len(colors)
	}
	return result
}

/*

1 - 1


*/
