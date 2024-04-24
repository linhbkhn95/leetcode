package numberofprovinces

func findCircleNum(isConnected [][]int) int {
	edges := make(map[int][]int, 0)
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[0]); j++ {
			if isConnected[i][j] == 1 {
				edges[i] = append(edges[i], j)
			}
		}
	}
	visted := make(map[int]bool, 0)
	result := 0
	for i := 0; i < len(isConnected); i++ {
		if visted[i] {
			continue
		}
		result++
		graph(edges, i, visted)
	}
	return result
}

func graph(edges map[int][]int, node int, visted map[int]bool) {
	if visted[node] {
		return
	}
	visted[node] = true
	neighbors, ok := edges[node]
	if !ok {
		return
	}
	for _, n := range neighbors {
		graph(edges, n, visted)
	}
}
