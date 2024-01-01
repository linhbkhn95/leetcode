package mergetwobinarytrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := TreeNode{
		Val: root1.Val + root2.Val,
	}
	root.Left = bfs(root1.Left, root2.Left)
	root.Right = bfs(root1.Right, root2.Right)
	return &root
}

func bfs(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		node := &TreeNode{
			Val: root2.Val,
		}
		node.Left = bfs(root2.Left, nil)
		node.Right = bfs(root2.Right, nil)
		return node
	}
	if root2 == nil {
		node := &TreeNode{
			Val: root1.Val,
		}
		node.Left = bfs(root1.Left, nil)
		node.Right = bfs(root1.Right, nil)
		return node
	}
	newNode := &TreeNode{
		Val: root1.Val + root2.Val,
	}

	newNode.Left = bfs(root1.Left, root2.Left)
	newNode.Right = bfs(root1.Right, root2.Right)
	return newNode
}
