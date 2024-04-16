package sumroottoleafnumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {

}

func dfs(root *TreeNode, result []int) (level int, addon int, sum int) {

	// leaf of tree
	if root.Left == nil && root.Right == nil {
		return 0, 0, root.Val
	}
	var count int
	if root.Left != nil {
		count += 1
		leftLevel, leftAddon, leftSum := dfs(root.Left, result)
		sum += leftSum
		addon += leftAddon
		if sum >= 10 {
			addon += 1
			sum -= 10
		}

	}
	if root.Right != nil {
		count += 1
		rightLevel, rightAddon, rightSum := dfs(root.Right, result)
		sum += rightSum
		if sum >= 10 {
			addon += 1
			sum -= 10
		}
	}
	return level, addon, sum
}

/*
0 - 1
5+1 = 6
9*2 = 18 => 8
0 -> 0
4+8 = 12 => 2
4*2 + 2 = 10
10 2 6


*/
