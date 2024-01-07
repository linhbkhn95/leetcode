package rotatelist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	t := head
	l := 0
	var tail *ListNode
	for t != nil {
		tail = t
		t = t.Next
		l++
	}
	k = k % l
	if k == 0 {
		return head
	}

	newHead := head
	var preNewHead *ListNode
	for i := 0; i < l-k; i++ {
		preNewHead = newHead
		newHead = newHead.Next
	}

	tail.Next = head
	preNewHead.Next = nil
	return newHead
}

func rotateRight1(head *ListNode, k int) *ListNode {
	t := head
	nodes := make([]*ListNode, 0)
	for t != nil {
		nodes = append(nodes, t)
		t = t.Next
	}
	l := len(nodes)
	k = k % l
	if k == 0 {
		return head
	}

	head = nodes[l-k]
	nodes[l-1].Next = nodes[0]
	nodes[l-k-1].Next = nil
	return head
}
