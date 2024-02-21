package main

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	minHeap := &MinHeap{}
	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			heap.Push(minHeap, diff)
			if minHeap.Len() > ladders {
				bricks -= heap.Pop(minHeap).(int)
				if bricks < 0 {
					return i - 1
				}
			}
		}
	}
	return len(heights) - 1
}

type MinHeap []int

// Len implements heap.Interface.
func (m MinHeap) Len() int {
	return len(m)
}

// Less implements heap.Interface.
func (m MinHeap) Less(i int, j int) bool {
	return m[i] < m[j]
}

// Pop implements heap.Interface.
func (m *MinHeap) Pop() any {
	item := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return item
}

// Push implements heap.Interface.
func (m *MinHeap) Push(x any) {
	(*m) = append((*m), x.(int))
}

// Swap implements heap.Interface.
func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = (*MinHeap)(nil)
