package recoverbinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var preNode, first, second *TreeNode

	inOrder(root, func(node *TreeNode) {
		if node == nil {
			return
		}
		if preNode != nil {
			if preNode != nil && preNode.Val > node.Val {
				first = preNode
			}
			if first != nil && preNode.Val > node.Val {
				second = node
			}
		}
		preNode = node
	})
	first.Val, second.Val = second.Val, first.Val
}

func inOrder(root *TreeNode, doFunc func(node *TreeNode)) {
	if root == nil {
		return
	}
	inOrder(root.Left, doFunc)
	doFunc(root)
	inOrder(root.Right, doFunc)
}

func recoverTree2(root *TreeNode) {
	var prev, first, second *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if prev != nil {
			if first == nil && prev.Val >= node.Val {
				first = prev
			}
			if first != nil && prev.Val >= node.Val {
				second = node
			}
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)

	first.Val, second.Val = second.Val, first.Val
}

func recoverTree1(root *TreeNode) {
	traversal(root)

}

func traversal(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		if root.Left.Val > root.Val {
			swap(root.Left, root)
			traversal(root)
		}
		maxNode := traversal(root.Left)
		if maxNode != nil {
			if maxNode.Val > root.Val {
				swap(maxNode, root)
			}
		}
	}
	if root.Right != nil {
		if root.Right.Val < root.Val {
			swap(root.Right, root)
			traversal(root)
		}
		leftNode := traversal(root.Right)
		if leftNode != nil {
			if leftNode.Val < root.Val {
				swap(leftNode, root)
				traversal(root)

			}
		}
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left.Val >= root.Right.Val {
		swap(root.Left, root.Right)
		traversal(root)
	}
	return root.Right
}

func swap(node1, node2 *TreeNode) {
	temp := node1.Val
	node1.Val = node2.Val
	node2.Val = temp
}
