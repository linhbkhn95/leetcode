package main

func main() {

}

func minJumps(arr []int) int {
	v := len(arr)
	g := NewGraph(v)
	mapping := make(map[int][]int, v)

	for i := 1; i < v-1; i++ {
		g.AddEdge(i, []int{i - 1, i + 1})
	}

	for i := 0; i < v; i++ {
		if _, ok := mapping[arr[i]]; ok {
			mapping[arr[i]] = append(mapping[arr[i]], i)
		} else {
			mapping[arr[i]] = []int{i}
		}
	}

	for _, v := range mapping {
		if len(v) > 1 {
			for i := 0; i < len(v); i++ {
				g.AddEdge(v[i], v[i:])
			}
		}
	}
	g.AddEdge(0, []int{1})
	g.AddEdge(v-1, []int{v - 2})
	return g.DFS()
}

type graph struct {
	adjs [][]int
	v    int
}

func NewGraph(v int) *graph {
	g := &graph{}
	adjs := make([][]int, v)
	for i := 1; i < v; i++ {
		adjs[i] = []int{}
	}
	g.adjs = adjs
	g.v = v
	return g
}

func (g *graph) AddEdge(position int, edges []int) {
	g.adjs[position] = append(g.adjs[position], edges...)
}

func (g *graph) DFSUtils(v int, visited []bool, currentLen int, minLen *int) {
	currentLen += 1
	if currentLen >= *minLen {
		return
	}
	visited[v] = true
	if v == g.v-1 {
		if currentLen < *minLen {
			*minLen = currentLen
		}
		return
	}
	for _, edge := range g.adjs[v] {
		if !visited[edge] || v != edge {
			g.DFSUtils(edge, visited, currentLen, minLen)
			if *minLen-currentLen == 1 {
				return
			}
		}
	}

}

func (g *graph) DFS() int {
	visited := make([]bool, g.v)

	for i := 0; i < g.v; i++ {
		visited[i] = false
	}
	var currentLen, minLen int
	minLen = g.v
	g.DFSUtils(0, visited, currentLen, &minLen)
	return minLen - 1
}
