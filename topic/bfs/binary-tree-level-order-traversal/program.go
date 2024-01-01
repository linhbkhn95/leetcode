package binarytreelevelordertraversal

import "sort"

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

func levelOrder(root *TreeNode) [][]int {
	bag := make([][]int, 0)
	bfs1(root, &bag, 0)
	return bag
}

func bfs1(root *TreeNode, bag *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*bag) == level {
		(*bag) = append((*bag), []int{})
	}
	(*bag)[level] = append((*bag)[level], root.Val)
	bfs1(root.Left, bag, level+1)
	bfs1(root.Right, bag, level+1)
}

func levelOrder1(root *TreeNode) [][]int {
	bag := make(map[int][]int, 0)
	bfs(root, bag, 0)
	type nodeLevel struct {
		Level    int
		Elements []int
	}
	arr := make([]*nodeLevel, 0, len(bag))
	for level, elements := range bag {
		arr = append(arr, &nodeLevel{
			Level:    level,
			Elements: elements,
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Level < arr[j].Level
	})
	result := make([][]int, len(arr))
	for i, nodeLevel := range arr {
		result[i] = nodeLevel.Elements
	}
	return result
}

func bfs(root *TreeNode, bag map[int][]int, level int) {
	if root == nil {
		return
	}
	_, ok := bag[level]
	if !ok {
		bag[level] = []int{root.Val}
	} else {
		bag[level] = append(bag[level], root.Val)
	}
	if root.Left != nil {
		bfs(root.Left, bag, level+1)
	}
	if root.Right != nil {
		bfs(root.Right, bag, level+1)
	}
}
