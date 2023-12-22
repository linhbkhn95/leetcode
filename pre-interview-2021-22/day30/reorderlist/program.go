package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

	h := head
	var arr []*ListNode
	for h != nil {
		arr = append(arr, h)
		h = h.Next
	}
	l := len(arr)
	i := l - 1
	k := 0
	tail := arr[i]
	t := head
	next := head.Next
	for t != nil && next != nil && tail != next && i > 0 && k < i {

		tail.Next = next
		t.Next = tail
		i -= 1
		k++
		tail = arr[i]
		t = next
		next = next.Next

	}
	if next != nil && tail != nil && next.Val == tail.Val {
		t.Next = next
		t = t.Next
	}

	t.Next = nil

}

func main() {
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					// Next: &ListNode{
					// 	Val: 5,
					// },
				},
			},
		},
	}
	reorderList(&node)
	display(&node)
}

func display(head *ListNode) {
	t := head
	count := 0
	for t != nil {
		if count > 6 {
			break
		}
		count++
		fmt.Println(t)
		t = t.Next
	}
}
