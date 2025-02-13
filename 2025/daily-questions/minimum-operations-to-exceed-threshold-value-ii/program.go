package minimumoperationstoexceedthresholdvalueii

import (
	"container/heap"
)

func minOperations(nums []int, k int) int {
	minHeap := &MinHead{}
	for _, num := range nums {
		heap.Push(minHeap, num)
	}
	result := 0
	var min1, min2, newNum int
	for minHeap.Len() > 1 {
		min1 = heap.Pop(minHeap).(int)
		if min1 >= k {
			return result
		}
		min2 = heap.Pop(minHeap).(int)
		newNum = min1*2 + min2
		heap.Push(minHeap, newNum)
		result += 1
	}
	min1 = heap.Pop(minHeap).(int)
	if min1 >= k {
		return result
	}
	return -1
}

type MinHead []int

// Len implements heap.Interface.
func (m *MinHead) Len() int {
	return len(*m)
}

// Less implements heap.Interface.
func (m MinHead) Less(i int, j int) bool {
	return m[i] < m[j]
}

// Pop implements heap.Interface.
func (m *MinHead) Pop() any {
	item := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return item
}

// Push implements heap.Interface.
func (m *MinHead) Push(x any) {
	(*m) = append((*m), x.(int))
}

// Swap implements heap.Interface.
func (m MinHead) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

var _ heap.Interface = (*MinHead)(nil)
