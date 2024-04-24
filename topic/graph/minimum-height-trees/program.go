package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n <= 1 {
		ans := make([]int, 0)
		for i := 0; i < n; i++ {
			ans = append(ans, i)
		}
		return ans
	}

	graph := make(map[int][]int)
	inDegree := make([]int, n)

	for _, edge := range edges {
		if _, ok := graph[edge[0]]; !ok {
			graph[edge[0]] = make([]int, 0)
		}

		if _, ok := graph[edge[1]]; !ok {
			graph[edge[1]] = make([]int, 0)
		}

		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		inDegree[edge[0]]++
		inDegree[edge[1]]++
	}

	queue := make([]int, 0)

	// find leaf nodes
	for i, degree := range inDegree {
		if degree == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		size := len(queue)
		n -= size

		for i := 0; i < size; i++ {
			inDegree[queue[i]]--

			for _, to := range graph[queue[i]] {
				inDegree[to]--

				// next round leaf node
				if inDegree[to] == 1 {
					queue = append(queue, to)
				}
			}
		}

		queue = queue[size:]
	}

	return queue
}
