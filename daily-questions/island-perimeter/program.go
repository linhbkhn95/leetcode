package islandperimeter

func islandPerimeter(grid [][]int) int {
	p := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				miniP := 4
				if i-1 >= 0 && grid[i-1][j] == 1 {
					miniP -= 1
				}
				if i+1 <= m-1 && grid[i+1][j] == 1 {
					miniP -= 1
				}
				if j-1 >= 0 && grid[i][j-1] == 1 {
					miniP -= 1
				}
				if j+1 <= n-1 && grid[i][j+1] == 1 {
					miniP -= 1
				}
				p += miniP
			}
		}
	}
	return p
}
