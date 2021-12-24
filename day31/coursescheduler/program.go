package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := NewGraph(numCourses)
	for _, pre := range prerequisites {
		graph.addEdge(pre[0], pre[1])
	}
	result := graph.DFS()
	if len(result) == numCourses {
		return result
	}
	return []int{}
}

type graph struct {
	adjs [][]int
	v    int
}

func NewGraph(v int) *graph {
	g := &graph{}
	g.v = v
	adjs := make([][]int, v)
	for i := 0; i < v; i++ {
		adjs[i] = []int{}
	}
	g.adjs = adjs
	return g
}

func (g *graph) addEdge(v, w int) {
	g.adjs[v] = append(g.adjs[v], w)
}

func (g *graph) DFSUtil(v int, visited []bool, flash []bool, bag *[]int) bool {
	if flash[v] {
		return false
	}
	flash[v] = true

	if len(g.adjs[v]) == 0 {
		visited[v] = true
	}
	for _, adj := range g.adjs[v] {
		k := adj
		if !visited[k] {
			if !g.DFSUtil(k, visited, flash, bag) {
				return false
			}
		}
	}
	visited[v] = true
	*bag = append(*bag, v)
	return true
}

func (g *graph) DFS() []int {
	visited := make([]bool, g.v)
	flash := make([]bool, g.v)
	for i := 0; i < g.v; i++ {
		visited[i] = false
		flash[i] = false
	}
	result := make([]int, 0, g.v)
	for i := 0; i < g.v; i++ {
		if !visited[i] {
			if !g.DFSUtil(i, visited, flash, &result) {
				return nil
			}
		}
	}
	return result
}

func main() {
	g := NewGraph(6)

	g.addEdge(1, 0)
	g.addEdge(2, 0)
	g.addEdge(3, 4)
	g.addEdge(3, 5)
	g.addEdge(4, 0)
	g.addEdge(2, 3)

	fmt.Println("Following is Depth First Traversal " +
		"(starting from vertex 2)")

	fmt.Println(g.DFS())
}
