package firstcompletelypaintedroworcolumn

func firstCompleteIndex(arr []int, mat [][]int) int {
	m, n := len(mat), len(mat[0])
	paintedRows := make([]int, m+1)
	paintedCols := make([]int, n+1)
	matInx := make([][2]int, m*n+1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matInx[mat[i][j]] = [2]int{i, j}
		}
	}
	for k, num := range arr {
		i, j := matInx[num][0], matInx[num][1]
		paintedRows[i] += 1
		paintedCols[j] += 1
		if paintedRows[i] == n || paintedCols[j] == m {
			return k
		}
	}
	return 0
}
