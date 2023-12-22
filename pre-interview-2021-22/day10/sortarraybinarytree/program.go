package sortarraybinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, start, end int) *TreeNode {

	if start >= end {
		return nil
	}
	avg := (start + end) / 2

	root := &TreeNode{
		Val: nums[avg],
	}
	root.Left = buildBST(nums, start, avg)
	root.Right = buildBST(nums, avg+1, end)
	return root
}
