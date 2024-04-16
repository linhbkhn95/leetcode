package main

func findMinHeightTrees(n int, edges [][]int) []int {
	nodes := make(map[int][]int, 0)
	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])
	}
	footprint := make(map[int]bool)
	hightTreeResult := make(map[int]int)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = dfs(nodes, n, i, footprint, hightTreeResult)
	}
	result := []int{}
	minV := n
	for i, v := range arr {
		if minV > v {
			minV = v
			result = []int{i}
		} else if minV == v {
			result = append(result, i)
		}
	}
	return result
}

func dfs(nodes map[int][]int, n, node int, footprint map[int]bool, hightTreeResult map[int]int) int {
	if footprint[node] {
		return n
	}
	nds, ok := nodes[node]
	if !ok {
		return 0
	}
	var minV int
	footprint[n] = true
	for _, nn := range nds {
		if v, ok := hightTreeResult[n]; ok {
			minV = min(minV, v+1)
			if 1+v > hightTreeResult[node] {
				hightTreeResult[node] = 1 + v
			}
			continue
		}
		minV = min(minV, 1+dfs(nodes, n, nn, footprint, hightTreeResult))
	}
	delete(footprint, node)
	hightTreeResult[node] = minV
	return minV
}
