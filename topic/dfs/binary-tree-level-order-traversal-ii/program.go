package binarytreelevelordertraversalii

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {

	bag := make(map[int][]int, 0)
	dfs(root, 0, bag)

	type Element struct {
		Level    int
		Elements []int
	}
	elements := make([]*Element, 0, len(bag))
	for level, e := range bag {
		elements = append(elements, &Element{
			Level:    level,
			Elements: e,
		})
	}
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Level > elements[j].Level
	})
	result := make([][]int, 0)
	for _, ele := range elements {
		result = append(result, ele.Elements)
	}
	return result
}

func dfs(root *TreeNode, level int, bag map[int][]int) {
	if root == nil {
		return
	}
	bag[level] = append(bag[level], root.Val)

	dfs(root.Left, level+1, bag)
	dfs(root.Right, level+1, bag)
}
