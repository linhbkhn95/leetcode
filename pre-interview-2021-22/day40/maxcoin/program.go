package maxcoin

import (
	"fmt"
	"sort"
)

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func maxCoins(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	index := make(map[int]*Node, l)
	temp := make([]int, l-2)
	head := &Node{
		Val: nums[0],
	}
	pre := head
	for i := 1; i < l-1; i++ {
		node := &Node{
			Val: nums[i],
		}
		pre.Next = node
		node.Pre = pre
		pre = node
		index[nums[i]] = node
		temp[i-1] = nums[i]
	}
	pre.Next = &Node{
		Val: nums[l-1],
		Pre: pre,
	}
	fmt.Println(index)
	sort.Ints(temp)
	total := 0

	for i := 0; i < len(temp); i++ {
		fmt.Println(temp[i])
		node := index[temp[i]]
		fmt.Println(node.Pre.Val, node.Val, node.Next.Val)

		total += node.Pre.Val * node.Val * node.Next.Val
		removeNode(node)
	}
	total += nums[l-1] * nums[0]
	if nums[0] > nums[l-1] {
		total += nums[0]
	} else {
		total += nums[l-1]
	}
	return total
}

func removeNode(node *Node) {
	pre := node.Pre
	next := node.Next
	next.Pre = pre
	pre.Next = next
}
