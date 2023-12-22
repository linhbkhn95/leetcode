package main

import (
	"container/heap"
	"fmt"
)

type SmallestInfiniteSet struct {
	removedNumbers map[int]struct{}
	heapMin        *HeapMin
}

func main() {
	s := Constructor()
	s.AddBack(2)
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	// s.AddBack(1)
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())

	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	s.AddBack(8)

	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
	fmt.Println(s.PopSmallest())
}

func Constructor() SmallestInfiniteSet {
	h := HeapMin{}
	for i := 1; i <= 1000; i++ {
		heap.Push(&h, i)
	}
	return SmallestInfiniteSet{
		removedNumbers: make(map[int]struct{}),
		heapMin:        &h,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	min := heap.Pop(this.heapMin)
	value := min.(int)
	this.removedNumbers[value] = struct{}{}
	return value
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.removedNumbers[num]; ok {
		heap.Push(this.heapMin, num)
		delete(this.removedNumbers, num)
	}
}

type HeapMin []int

// Len implements heap.Interface.
func (h HeapMin) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h HeapMin) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Pop implements heap.Interface.
func (h *HeapMin) Pop() any {
	i := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	// heap.Fix(h, n-2)
	return i
}

// Push implements heap.Interface.
func (h *HeapMin) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements heap.Interface.
func (h HeapMin) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

var _ heap.Interface = (*HeapMin)(nil)

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
