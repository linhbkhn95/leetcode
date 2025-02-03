package gridgame

import "container/heap"

func gridGame(grid [][]int) int64 {
	r, c := 2, len(grid[0])
	visited := make([][]bool, 2)
	for i := 0; i < 2; i++ {
		visited[i] = make([]bool, c)
	}
	dp(grid, r, c, 0, 0, visited)
	visited[0][0] = false
	return dp(grid, r, c, 0, 0, visited) - int64(grid[1][c-1])
}

func dp(grid [][]int, r, c, i, j int, visited [][]bool) int64 {
	if i == r || j == c {
		return 0
	}
	if visited[i][j] {
		return 0
	}
	var down, right int64 = 0, 0
	if i == 0 {
		down = dp(grid, r, c, i+1, j, visited)
	}
	if j < c-1 {
		right = dp(grid, r, c, i, j+1, visited)
	}
	if down > right && down > 0 {
		visited[i+1][j] = true
	} else if right > 0 {
		visited[i][j+1] = true
	}
	return int64(grid[i][j]) + max(down, right)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func dijkstra(matrix [2][]int, r, c int, root [2]int, distances [2][]int, path [2][][2]int) {
	minHeap := &MaxHeap{}

	heap.Push(minHeap, &Element{
		Node:     root,
		Distance: 0,
	})
	distances[root[0]][root[1]] = 0
	for minHeap.Len() > 0 {
		e := heap.Pop(minHeap).(*Element)
		w := e.Distance
		node := e.Node
		neighbors := [][2]int{}
		if node[1] < c {
			neighbors = append(neighbors, [2]int{node[0], node[1] + 1})
		}
		if node[0] < 1 {
			neighbors = append(neighbors, [2]int{node[0] + 1, node[1]})
		}
		for _, neighbor := range neighbors {
			if w+matrix[neighbor[0]][neighbor[1]] > distances[neighbor[0]][neighbor[1]] {
				distances[neighbor[0]][neighbor[1]] = w + matrix[neighbor[0]][neighbor[1]]
				heap.Push(minHeap, &Element{
					Node:     neighbor,
					Distance: distances[neighbor[0]][neighbor[1]],
				})
				path[neighbor[0]][neighbor[1]] = e.Node
			}
		}
	}
}

type Element struct {
	Node     [2]int
	Distance int
}

type MaxHeap []*Element

// Len implements heap.Interface.
func (m MaxHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MaxHeap) Less(i int, j int) bool {
	return m[i].Distance < m[j].Distance
}

// Pop implements heap.Interface.
func (m *MaxHeap) Pop() any {
	item := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return item
}

// Push implements heap.Interface.
func (m *MaxHeap) Push(x any) {
	(*m) = append((*m), x.(*Element))
}

// Swap implements heap.Interface.
func (m MaxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = (*MaxHeap)(nil)
