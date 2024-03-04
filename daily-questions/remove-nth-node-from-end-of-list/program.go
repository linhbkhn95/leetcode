package removenthnodefromendoflist

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Val: -1, Next: head}

	prev, cur := temp, temp

	for cur.Next != nil {
		if n <= 0 {
			prev = prev.Next
		}

		cur = cur.Next
		n--
	}

	target := prev.Next
	prev.Next = target.Next

	return temp.Next
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	list := make([]*ListNode, 0)
// 	node := head
// 	for node != nil {
// 		list = append(list, node)
// 		node = node.Next
// 	}
// 	l := len(list)
// 	properNode := list[l-n]
// 	if n == l {
// 		head = head.Next
// 	} else {
// 		list[l-n-1].Next = properNode.Next
// 	}
// 	return head
// }
