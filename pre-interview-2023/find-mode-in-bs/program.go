package findmodeinbs

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

func findMode(root *TreeNode) []int {
	bag := make(map[int]int, 0)
	Move(root, bag)
	max := 0
	result := make([]int, 0)

	for k, v := range bag {
		if v > max {
			max = v
			result = append(result[:0], k)
		} else if v == max {
			result = append(result, k)
		}
	}

	return result
}

func Move(root *TreeNode, bag map[int]int) {
	if root == nil {
		return
	}
	bag[root.Val]++
	if root.Left != nil {
		Move(root.Left, bag)
	}
	if root.Right != nil {
		Move(root.Right, bag)
	}
}
