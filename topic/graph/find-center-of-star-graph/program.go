package findcenterofstargraph

func findCenter(edges [][]int) int {
	count := make(map[int]int)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		if count[x] == 1 {
			return x
		}
		if count[y] == 1 {
			return y
		}
		count[x]++
		count[y]++
	}
	return -1
}
