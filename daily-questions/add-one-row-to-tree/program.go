package addonerowtotree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newRoot := &TreeNode{
			Val:  val,
			Left: root,
		}
		return newRoot
	}
	dfs(root, val, depth, 1)
	return root
}

func dfs(root *TreeNode, val int, depth int, currentDepth int) {
	if root == nil {
		return
	}
	if currentDepth == depth-1 {
		leftRoot := root.Left
		rightRoot := root.Right
		newLeft := &TreeNode{
			Val: val,
		}
		newRight := &TreeNode{
			Val: val,
		}
		root.Left = newLeft
		root.Right = newRight
		newLeft.Left = leftRoot
		newRight.Right = rightRoot
		return
	}
	dfs(root.Left, val, depth, currentDepth+1)
	dfs(root.Right, val, depth, currentDepth+1)
}

/*


 */
