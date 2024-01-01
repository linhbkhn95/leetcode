package minimumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	leftDeth := 1 + minDepth(root.Left)
	rightDeth := 1 + minDepth(root.Right)

	if leftDeth < rightDeth {
		return leftDeth
	}
	return rightDeth
}
