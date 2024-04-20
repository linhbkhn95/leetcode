package findallgroupsoffarmland

func findFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	result := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] == 1 {
				result = append(result, extractCandicate(land, m, n, i, j))
			}
		}
	}
	return result
}

func extractCandicate(land [][]int, m, n, i, j int) []int {
	result := make([]int, 0, 4)
	// find top left
	for i >= 0 && j >= 0 {
		if i >= 1 && land[i-1][j] == 1 {
			i--
			continue
		}
		if j >= 1 && land[i][j-1] == 1 {
			j--
			continue
		}
		break
	}
	// find bottom right
	result = append(result, []int{i, j}...)
	for i < m && j < n {
		if i < m-1 && land[i+1][j] == 1 {
			i++
			continue
		}
		if j < n-1 && land[i][j+1] == 1 {
			j++
			continue
		}
		break
	}
	result = append(result, []int{i, j}...)

	// mark crossed
	for k := result[0]; k <= result[2]; k++ {
		for o := result[1]; o <= result[3]; o++ {
			land[k][o] = -1
		}
	}
	return result
}
