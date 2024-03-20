package merge_in_between_linked_lists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var preNodeA, preNode, nextNodeB *ListNode
	head := list1
	index := 0
	for head != nil {
		if index == a {
			preNodeA = preNode
		}
		if index == b {
			nextNodeB = head.Next
			break
		}
		preNode = head
		head = head.Next
		index++
	}
	head = list2
	for head.Next != nil {
		head = head.Next
	}
	head.Next = nextNodeB

	if preNodeA != nil {
		preNodeA.Next = list2
		return list1
	}

	return list2
}

/*

Input: list1 = [10,1,13,6,9,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
Output: [10,1,13,1000000,1000001,1000002,5]
Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place. The blue edges and nodes in the above figure indicate the result.
Example 2:


Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
Explanation: The blue edges and nodes in the above figure indicate the result.

*/
