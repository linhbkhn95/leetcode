package reorder_list

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

func reorderList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	count := 0
	h := head
	for h != nil {
		count++
		h = h.Next
	}
	h = head
	index := 1
	var right, tail, nextRight *ListNode
	for h != nil {
		if count%2 == 1 && index == (count+1)/2 || count%2 == 0 && index == count/2 {
			right = h.Next
			for right != nil {
				nextRight = right.Next
				right.Next = tail
				tail = right
				right = nextRight
			}
			break
		}
		index++
		h = h.Next
	}
	h = head
	var nextHead, nextTail *ListNode
	for tail != nil {
		nextHead = h.Next
		nextTail = tail.Next
		h.Next = tail
		tail.Next = nextHead
		tail = nextTail
		h = nextHead
	}
	if h != nil {
		h.Next = nil
	}
	return head
}

// 1->2->3->4->5 -> 1-5->2->4->3
