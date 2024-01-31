package triangle

type key struct {
	i, j int
}

func minimumTotal(triangle [][]int) int {
	return dp(triangle, len(triangle), 0, 0, map[key]int{})
}

func dp(triangle [][]int, m, i, j int, footprint map[key]int) int {
	if i == m-1 {
		return triangle[i][j]
	}
	k := key{
		i: i,
		j: j,
	}
	if v, ok := footprint[k]; ok {
		return v
	}
	minV := triangle[i][j] + min(dp(triangle, m, i+1, j, footprint), dp(triangle, m, i+1, j+1, footprint))
	footprint[k] = minV
	return minV
}
