package minimum_absolute_difference_in_bst

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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
func getMinimumDifference(root *TreeNode) int {
	result := math.MaxInt
	move(root, &result)
	return result
}

func move(root *TreeNode, minValue *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		*minValue = min(*minValue, root.Val-root.Left.Val)
		*minValue = min(*minValue, root.Val-MaxRight(root.Left))
		move(root.Left, minValue)
	}
	if root.Right != nil {
		*minValue = min(*minValue, root.Right.Val-root.Val)
		*minValue = min(*minValue, MinLeft(root.Right)-root.Val)
		move(root.Right, minValue)
	}
}

func MinLeft(root *TreeNode) int {
	if root == nil {
		return 0
	}
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func MaxRight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}
