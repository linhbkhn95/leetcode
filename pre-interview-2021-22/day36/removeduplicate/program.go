package removeduplicate

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	var pre *ListNode
	var ppre *ListNode
	var result *ListNode
	for head != nil {
		flag := false
		if pre != nil && pre.Val == head.Val {
			pre = ppre
			if ppre != nil {
				ppre.Next = nil
			}
		}
		for head != nil && head.Next != nil && head.Val == head.Next.Val {

			flag = true
			head = head.Next
		}
		if flag && head != nil {
			head = head.Next
		}
		if pre == nil {
			result = head
		}
		if pre != ppre {
			ppre = pre
		}
		pre = head
		if ppre != nil {
			ppre.Next = pre
		}
		if !flag && head != nil {
			head = head.Next
		}
	}
	return result
}

func ThreeSum(arr []int, k int) [][]int {
	sort.Ints(arr)
	l := len(arr)
	result := [][]int{}
	if l < 3 {
		return nil
	}
	i := 0
	for i < l {
		if i > 0 && arr[i] == arr[i-1] {
			i++
			continue
		}
		j := i + 1
		p := l - 1
		for j < p {

			if arr[i]+arr[j]+arr[p] == k {
				result = append(result, []int{arr[i], arr[j], arr[p]})
				h := j
				for h < l && arr[h] == arr[j] {
					h++
				}
				j = h

			} else if arr[i]+arr[j]+arr[p] > k {
				p--
			} else {
				j++
			}
		}
		i++
	}
	return result
}
