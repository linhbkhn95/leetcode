package leafsimilartrees

import (
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return reflect.DeepEqual(dfs(root1), dfs(root2))
}
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return append(dfs(root.Left), dfs(root.Right)...)
}
func leafSimilar1(root1 *TreeNode, root2 *TreeNode) bool {
	nodes := make(map[int]int, 0)
	order1 := 1
	dfs1(root1, nodes, &order1, false)
	order2 := 1
	dfs1(root2, nodes, &order2, true)
	if order1 != order2 {
		return false
	}
	return len(nodes) == 0
}

func dfs1(root *TreeNode, nodes map[int]int, order *int, shouldDelete bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if shouldDelete {
			num, ok := nodes[*order]
			if ok && num == root.Val {
				delete(nodes, *order)
			}
		} else {
			nodes[*order] = root.Val
		}
		*order++
	}
	dfs1(root.Left, nodes, order, shouldDelete)
	dfs1(root.Right, nodes, order, shouldDelete)
}
