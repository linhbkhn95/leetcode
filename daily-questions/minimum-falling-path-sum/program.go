package minimumfallingpathsum

import "math"

type key struct {
	i int
	j int
}

func minFallingPathSum(matrix [][]int) int {
	footprint := make(map[key]int)
	minResult := math.MaxInt32
	m, n := len(matrix), len(matrix[0])
	for j := 0; j < len(matrix[0]); j++ {
		minResult = min(minResult, dp(matrix, m, n, 0, j, footprint))
	}
	return minResult
}

func dpWithoutRecursive(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
		copy(dp[i], matrix[i])
	}
	m, n := len(matrix), len(matrix[0])
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			minDP := math.MaxInt32
			if j == 0 {
				minDP = min(dp[i+1][j], dp[i+1][j+1])
			} else if j == n-1 {
				minDP = min(dp[i+1][j], dp[i+1][j-1])

			} else {
				minDP = min(dp[i+1][j], dp[i+1][j+1])
				minDP = min(minDP, dp[i+1][j-1])
			}
			dp[i][j] = matrix[i][j] + minDP
		}
	}
	minResult := math.MaxInt32
	for i := 0; i < n; i++ {
		minResult = min(minResult, dp[0][i])
	}
	return minResult
}

func dp(matrix [][]int, m, n, i, j int, footprint map[key]int) int {
	k := key{
		i,
		j,
	}
	if i == n || j == m || i < 0 || j < 0 {
		return math.MaxInt32
	}
	if i == m-1 {
		return matrix[i][j]
	}
	if v, ok := footprint[k]; ok {
		return v
	}

	minResult := math.MaxInt32
	minResult = min(minResult, matrix[i][j]+dp(matrix, m, n, i+1, j-1, footprint))
	minResult = min(minResult, matrix[i][j]+dp(matrix, m, n, i+1, j, footprint))
	minResult = min(minResult, matrix[i][j]+dp(matrix, m, n, i+1, j+1, footprint))

	footprint[k] = minResult
	return minResult
}
