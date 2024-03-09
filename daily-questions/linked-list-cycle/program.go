package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var pre *ListNode
	for head != nil {
		if head.Next == head {
			return true
		}
		pre = head
		head = head.Next
		pre.Next = pre
	}
	return false
}
