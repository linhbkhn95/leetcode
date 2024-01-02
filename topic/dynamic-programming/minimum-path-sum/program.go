package minimumpathsum

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	footprint := make(map[key]int)
	var s int
	if m == 1 {
		for _, n := range grid[0] {
			s += n
		}
		return s
	}
	if n == 1 {
		for _, g := range grid {
			s += g[0]
		}
		return s
	}
	return min(grid[0][0]+dp(grid, m, n, 0, 1, footprint), grid[0][0]+dp(grid, m, n, 1, 0, footprint))
}

type key struct {
	i int
	j int
}

func dp(grid [][]int, m, n, i, j int, footprint map[key]int) int {

	if i == m-1 && j == n-1 {
		return grid[i][j]
	}
	k := key{
		i: i,
		j: j,
	}
	sum, ok := footprint[k]
	if ok {
		return sum
	}
	if i == m-1 {
		return grid[i][j] + dp(grid, m, n, i, j+1, footprint)
	}
	if j == n-1 {
		return grid[i][j] + dp(grid, m, n, i+1, j, footprint)
	}
	minX := min(grid[i][j]+dp(grid, m, n, i+1, j, footprint), grid[i][j]+dp(grid, m, n, i, j+1, footprint))
	footprint[k] = minX
	return minX

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
