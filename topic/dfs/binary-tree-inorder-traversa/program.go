package binarytreeinordertraversa

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

func inorderTraversal(root *TreeNode) []int {
	bag := make([]int, 0)
	traversal(root, &bag)
	return bag
}

func traversal(root *TreeNode, bag *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, bag)
	(*bag) = append((*bag), root.Val)
	traversal(root.Right, bag)
}
