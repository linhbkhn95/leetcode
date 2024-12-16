package finalarraystateafterkmultiplicationoperationsi

import (
	"container/heap"
)

type Element struct {
	Offset int
	Value  int
}

func getFinalState(nums []int, k int, multiplier int) []int {
	h := &Elements{}
	for i, num := range nums {
		*h = append(*h, &Element{
			Offset: i,
			Value:  num,
		})
	}
	heap.Init(h)
	for i := 1; i <= k; i++ {
		element := heap.Pop(h)
		e := element.(*Element)
		e.Value *= multiplier
		heap.Push(h, element)
		heap.Fix(h, h.Len()-1)
	}
	for _, i := range *h {
		nums[i.Offset] = i.Value
	}
	return nums
}

type Elements []*Element

// Pop implements heap.Interface.
func (e *Elements) Pop() any {
	item := (*e)[e.Len()-1]
	(*e) = (*e)[:e.Len()-1]
	return item
}

// Push implements heap.Interface.
func (e *Elements) Push(x any) {
	item := x.(*Element)
	(*e) = append((*e), item)
}

// Len implements sort.Interface.
func (e Elements) Len() int {
	return len(e)
}

// Less implements sort.Interface.
func (e Elements) Less(i int, j int) bool {
	if e[i].Value < e[j].Value {
		return true
	}
	if e[i].Value == e[j].Value {
		return e[i].Offset < e[j].Offset
	}
	return false
}

// Swap implements sort.Interface.
func (e Elements) Swap(i int, j int) {
	e[i], e[j] = e[j], e[i]
}

var _ heap.Interface = (*Elements)(nil)
