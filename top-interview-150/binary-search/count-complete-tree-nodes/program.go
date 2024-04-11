package countcompletetreenodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeLeft, nodeRight := 0, 0
	if root.Left != nil {
		nodeLeft = 1 + dfs(root.Left)
	}
	if root.Right != nil {
		nodeRight = 1 + dfs(root.Right)
	}
	return nodeLeft + nodeRight
}
