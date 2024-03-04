package findbottomlefttreevalue

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type Bag struct {
	Value int
	High  int
}

func findBottomLeftValue(root *TreeNode) int {
	result := Bag{
		Value: root.Val,
		High:  0,
	}
	dfs(root, &result, 0, false)
	return result.Value
}

func dfs(root *TreeNode, result *Bag, high int, isNodeLeft bool) {
	if root.Left == nil && root.Right == nil {
		if isNodeLeft {
			if result.High < high {
				result.Value = root.Val
				result.High = high
			}
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, result, high+1, true)
	}
	if root.Right != nil {
		dfs(root.Right, result, high+1, false)
	}
}
