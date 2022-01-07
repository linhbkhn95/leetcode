package linkedrandom

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	list        *ListNode
	lenList     int
	mappingList map[int]int
}

func Constructor(head *ListNode) Solution {
	mapping := make(map[int]int)

	t := head
	result := 0
	for t != nil {
		mapping[result+1] = t.Val
		t = t.Next
		result++
	}

	return Solution{
		list:        head,
		lenList:     result,
		mappingList: mapping,
	}
}

func (this *Solution) GetRandom() int {
	offset := rand.Intn(this.lenList)
	return this.mappingList[offset]
}
