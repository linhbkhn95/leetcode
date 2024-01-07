package partitionlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	t := head
	var pre, preGreateNode, greateNode *ListNode
	for t != nil {
		if t.Val >= x && greateNode == nil {
			preGreateNode = pre
			greateNode = t
		}
		if t.Val < x && greateNode != nil {
			if pre != nil {
				pre.Next = t.Next
			} else {
				head = t
			}
			if preGreateNode != nil {
				preGreateNode.Next = t
			} else {
				head = t
			}
			t.Next = greateNode

			preGreateNode = t
		}
		pre = t
		t = t.Next
	}
	return head
}
