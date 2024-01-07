package reverselinkedlistii

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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	t := head
	var preLeft *ListNode
	for i := 1; i < left; i++ {
		preLeft = t
		t = t.Next
	}
	var next, pre *ListNode

	for j := left; j <= right; j++ {
		if t == nil {
			break
		}
		next = t.Next

		t.Next = pre
		pre = t
		t = next
	}
	if preLeft != nil {
		preLeft.Next.Next = next
		preLeft.Next = pre
	} else {
		if next != nil {
			head.Next = next
		}
		head = pre

	}

	return head
}
