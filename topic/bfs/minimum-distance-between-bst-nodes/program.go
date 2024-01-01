package minimumdistancebetweenbstnodes

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

func minDiffInBST1(root *TreeNode) int {
	if root == nil {
		return math.MaxInt32
	}
	if root.Left == nil && root.Right == nil {
		return math.MaxInt32

	}
	if root.Right == nil {
		return min(root.Val-root.Left.Val, minDiffInBST(root.Left))
	}
	if root.Left == nil {
		return min(root.Right.Val-root.Val, minDiffInBST(root.Right))
	}
	minX := min(root.Val-root.Left.Val, root.Right.Val-root.Val)
	minX = min(minX, minDiffInBST(root.Left))
	minX = min(minX, minDiffInBST(root.Right))
	return minX
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return math.MaxInt32
	}
	if root.Left == nil && root.Right == nil {
		return math.MaxInt32

	}
	minX := math.MaxInt32
	if root.Left != nil {
		minX = min(root.Val-getMaxRight(root.Left), minX)
		minX = min(minX, minDiffInBST(root.Left))
	}
	if root.Right != nil {
		minX = min(minX, getMinLeft(root.Right)-root.Val)
		minX = min(minX, minDiffInBST(root.Right))
	}
	return minX
}

func getMaxRight(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}
	return getMaxRight(root.Right)
}

func getMinLeft(root *TreeNode) int {
	if root.Left == nil {
		return root.Val
	}
	return getMinLeft(root.Left)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
