package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)
	
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	right, left := root.Right, root.Left
	var maxNode *TreeNode
	if left != nil {
		root.Right = left
		maxNode = dfs(left)
	}
	if right != nil {
		if maxNode != nil {
			maxNode.Right = right
		}
		maxNode = dfs(right)
	}
	root.Left = nil
	return maxNode
}
