package findifpathexistsingraph

func validPath(n int, edges [][]int, source int, destination int) bool {
	nodes := make(map[int][]int, 0)
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}
	return graph(nodes, source, destination, map[int]bool{})
}

func graph(nodes map[int][]int, source, destination int, footprint map[int]bool) bool {
	if footprint[source] {
		return false
	}
	if source == destination {
		return true
	}
	candidates, ok := nodes[source]
	if !ok {
		return false
	}
	footprint[source] = true
	for _, can := range candidates {
		if graph(nodes, can, destination, footprint) {
			return true
		}
	}
	return false
}
