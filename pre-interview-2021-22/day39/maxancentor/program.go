package maxancentor

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return getMax(root, []int{})
}

func getMax(root *TreeNode, ancestorValues []int) int {
	ancestorValues = append(ancestorValues, root.Val)
	absLeft, absRight := 0, 0
	if root.Left != nil {
		absLeft = max(getMax(root.Left, ancestorValues), extractMax(root.Left.Val, ancestorValues))
	}

	if root.Right != nil {
		absRight = max(getMax(root.Right, ancestorValues), extractMax(root.Right.Val, ancestorValues))
	}
	ancestorValues = ancestorValues[:len(ancestorValues)-1]
	return max(absLeft, absRight)
}

func extractMax(nodeValue int, ancestorValues []int) int {
	max := 0
	for _, ancestorValue := range ancestorValues {
		abs := math.Abs(float64(ancestorValue) - float64(nodeValue))
		if abs > float64(max) {
			max = int(abs)
		}
	}

	return max
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}
