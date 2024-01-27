package outofboundarypaths

import "math"

var pw = int64(math.Pow10(9)) + 7

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	count := dp(m, n, maxMove, startRow, startColumn, map[key]int64{})
	return int(count)
}

type key struct {
	maxMove     int
	startRow    int
	startColumn int
}

func dp(m int, n int, maxMove int, startRow int, startColumn int, footprint map[key]int64) int64 {
	if startColumn < 0 || startColumn == n || startRow < 0 || startRow == m {
		return 1
	}
	if startColumn >= maxMove && n-1-startColumn >= maxMove && startRow >= maxMove && m-1-startRow >= maxMove {
		return 0
	}
	if maxMove == 0 {
		return 0
	}
	k := key{
		maxMove:     maxMove,
		startRow:    startRow,
		startColumn: startColumn,
	}
	if v, ok := footprint[k]; ok {
		return v
	}
	count := dp(m, n, maxMove-1, startRow-1, startColumn, footprint) + dp(m, n, maxMove-1, startRow+1, startColumn, footprint) + dp(m, n, maxMove-1, startRow, startColumn-1, footprint) + dp(m, n, maxMove-1, startRow, startColumn+1, footprint)
	count = count % pw
	footprint[k] = count
	return count
}
