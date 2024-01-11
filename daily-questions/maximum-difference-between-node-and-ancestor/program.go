package maximumdifferencebetweennodeandancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return getMaxDistance(root, root.Val, root.Val)
}

func getMaxDistance(root *TreeNode, low, high int) int {
	if root == nil {
		return abs(low - high)
	}
	low = min(root.Val, low)
	high = max(root.Val, high)
	left := getMaxDistance(root.Left, low, high)
	right := getMaxDistance(root.Right, low, high)
	return max(left, right)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
