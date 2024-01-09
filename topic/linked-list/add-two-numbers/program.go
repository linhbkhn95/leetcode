package addtwonumbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, newHead *ListNode
	var redundant int
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + redundant
		if sum >= 10 {
			redundant = 1
			sum = sum - 10
		} else {
			redundant = 0
		}
		node := &ListNode{
			Val: sum,
		}
		if head == nil {
			newHead = node
		} else {
			head.Next = node
		}
		head = node

		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + redundant
		if sum >= 10 {
			redundant = 1
			sum = sum - 10
		} else {
			redundant = 0
		}
		node := &ListNode{
			Val: sum,
		}

		if head == nil {
			newHead = node
		} else {
			head.Next = node
		}
		head.Next = node
		head = node
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + redundant
		if sum >= 10 {
			redundant = 1
			sum = sum - 10
		} else {
			redundant = 0
		}
		node := &ListNode{
			Val: sum,
		}
		if head == nil {
			newHead = node
		} else {
			head.Next = node
		}
		head = node
		l2 = l2.Next
	}
	if redundant > 0 {
		head.Next = &ListNode{
			Val: redundant,
		}
	}

	return newHead

}
