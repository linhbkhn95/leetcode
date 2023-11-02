package countnodesequaltoaverage

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	result := 0
	_, _, ok := calc(root, &result)
	if ok {
		result += 1
	}
	return result

}

func calc(root *TreeNode, result *int) (int, int, bool) {
	if root == nil {
		return 0, 0, false
	}
	total, count := root.Val, 1
	if root.Left != nil {
		sumLeft, countLeftNodes, ok := calc(root.Left, result)
		count += countLeftNodes
		total += sumLeft
		if ok {
			*result += 1
		}
	}
	if root.Right != nil {
		sumRight, countRightNodes, ok := calc(root.Right, result)
		total += sumRight
		count += countRightNodes
		if ok {
			*result += 1
		}
	}
	if total/count == root.Val {
		return total, count, true
	}
	return total, count, false
}
