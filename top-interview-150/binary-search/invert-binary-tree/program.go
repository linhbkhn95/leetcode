package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	dfs(root)
	return root
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	dfs(root.Left)
	dfs(root.Right)
}
