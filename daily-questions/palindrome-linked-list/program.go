package palindrome_linked_list

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

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	var count int
	h := head
	for h != nil {
		count++
		h = h.Next
	}
	h = head
	var pre, next *ListNode

	index := 0
	for h != nil {
		next = h.Next
		if count%2 == 1 && index == (count-1)/2 {
			return checkIsPalindrome(pre, next)
		}
		if count%2 == 0 && index == count/2 {
			return checkIsPalindrome(pre, h)
		}
		h.Next = pre
		pre = h
		h = next
		index++
	}
	return false
}

func checkIsPalindrome(left, right *ListNode) bool {
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	if left == nil && right == nil {
		return true
	}
	return false
}
