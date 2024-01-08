package rangesumofbst

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

func rangeSumBST(root *TreeNode, low int, high int) int {
	result := 0
	BTS(root, low, high, &result)
	return result
}

func BTS(root *TreeNode, low, hight int, result *int) {
	if root == nil {
		return
	}
	if root.Val >= low && root.Val <= hight {
		*result += root.Val
	}

	if root.Left != nil {
		BTS(root.Left, low, hight, result)
	}
	if root.Right != nil {
		BTS(root.Right, low, hight, result)
	}
}
