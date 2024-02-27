package diameterofbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	maxLength(root, &result)
	return result
}

func maxLength(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	var lenLeft, lenRight int
	if root.Left != nil {
		lenLeft = 1 + maxLength(root.Left, result)
	}
	if root.Right != nil {
		lenRight = 1 + maxLength(root.Right, result)
	}
	*result = max(lenLeft+lenRight, *result)
	return max(lenLeft, lenRight)
}
