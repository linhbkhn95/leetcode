package balancedbinarytree

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

func isBalanced(root *TreeNode) bool {
	_, ok := traversal(root)
	return ok
}

func traversal(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	if root.Left == nil && root.Right == nil {
		return 0, true
	}
	var leftDepth, rightDepth int
	if root.Left != nil {
		lDepth, ok := traversal(root.Left)
		if !ok {
			return 0, false
		}
		leftDepth = 1 + lDepth
	}
	if root.Right != nil {
		lRight, ok := traversal(root.Right)
		if !ok {
			return 0, false
		}
		rightDepth = 1 + lRight
	}
	if leftDepth-rightDepth >= 2 || leftDepth-rightDepth <= -2 {
		return 0, false
	}
	if leftDepth > rightDepth {
		return leftDepth, true
	}
	return rightDepth, true
}
