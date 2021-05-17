package longestdistinctpath

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type footprint map[int]bool

func Solution(root *TreeNode) int {
	fp := make(footprint)
	return solution(root, fp)
}

func solution(root *TreeNode, fp footprint) int {
	if root == nil {
		return 0
	}
	if _, ok := fp[root.Val]; ok {
		return 0
	}
	fp[root.Val] = true
	numberOfLeft := 1 + solution(root.Left, fp)
	numberOfRight := 1 + solution(root.Right, fp)
	delete(fp, root.Val)
	return max(numberOfLeft, numberOfRight)
}

func max(firstNumber, secondNumber int) int {
	if firstNumber > secondNumber {
		return firstNumber
	}
	return secondNumber
}
