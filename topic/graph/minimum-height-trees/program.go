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
		footprint[i] = true
		arr[i] = dfs(nodes, i, footprint, hightTreeResult)
		delete(footprint, i)

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

func dfs(nodes map[int][]int, node int, footprint map[int]bool, hightTreeResult map[int]int) int {

	nds, ok := nodes[node]
	if !ok {
		return 0
	}
	var maxV int
	for _, n := range nds {

		if v, ok := hightTreeResult[n]; ok {
			maxV = max(maxV, v+1)
			if 1+v > hightTreeResult[node] {
				hightTreeResult[node] = 1 + v
			}
			continue
		}
		if footprint[n] {
			continue
		}
		footprint[n] = true
		maxV = max(maxV, 1+dfs(nodes, n, footprint, hightTreeResult))
		delete(footprint, n)
	}
	hightTreeResult[node] = maxV
	return maxV
}
