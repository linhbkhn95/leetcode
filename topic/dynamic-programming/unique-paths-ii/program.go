package uniquepathsii

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	return dp(obstacleGrid, len(obstacleGrid), len(obstacleGrid[0]), 0, 0, make(map[key]int))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	obstacleGrid[0][0] = -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if (i+1) < m && obstacleGrid[i+1][j] != 1 {
				obstacleGrid[i+1][j] += obstacleGrid[i][j]
			}
			if (j+1) < n && obstacleGrid[i][j+1] != 1 {
				obstacleGrid[i][j+1] += obstacleGrid[i][j]
			}
		}
	}
	return -1 * obstacleGrid[m-1][n-1]
}

type key struct {
	i int
	j int
}

func dp(obstacleGrid [][]int, m, n, i, j int, footprint map[key]int) int {
	if i == m || j == n || obstacleGrid[i][j] == 1 {
		return 0
	}
	if i == m-1 && j == n-1 {
		return 1
	}
	k := key{
		i: i,
		j: j,
	}
	sum, ok := footprint[k]
	if ok {
		return sum
	}
	downWay := dp(obstacleGrid, m, n, i+1, j, footprint)
	righWay := dp(obstacleGrid, m, n, i, j+1, footprint)
	footprint[k] = downWay + righWay
	return downWay + righWay
}
