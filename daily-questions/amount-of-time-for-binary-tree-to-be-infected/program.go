package amountoftimeforbinarytreetobeinfected

import (
	"math"
)

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

type key struct {
	level  int
	isLeft bool
}

func amountOfTime(root *TreeNode, start int) int {
	maxDistance := 0
	nodes := make(map[int][]int)
	footprint := make(map[int]struct{})
	footprint[start] = struct{}{}
	bfs(root, nodes)
	backtracking(nodes, footprint, start, 0, &maxDistance)
	return maxDistance
}

func backtracking(nodes map[int][]int, footprint map[int]struct{}, start, currentDistance int, maxDistance *int) {

	edges, ok := nodes[start]
	if !ok {
		return
	}
	for _, edge := range edges {
		if _, ok := footprint[edge]; ok {
			continue
		}
		footprint[edge] = struct{}{}
		backtracking(nodes, footprint, edge, currentDistance+1, maxDistance)
		delete(footprint, edge)
	}
	*maxDistance = max(currentDistance, *maxDistance)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bfs(root *TreeNode, nodes map[int][]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		nodes[root.Val] = append(nodes[root.Val], root.Left.Val)
		nodes[root.Left.Val] = append(nodes[root.Left.Val], root.Val)
		bfs(root.Left, nodes)
	}

	if root.Right != nil {
		nodes[root.Val] = append(nodes[root.Val], root.Right.Val)
		nodes[root.Right.Val] = append(nodes[root.Right.Val], root.Val)
		bfs(root.Right, nodes)
	}
}

func amountOfTime1(root *TreeNode, start int) int {
	nodes := make(map[int]key)
	if root == nil {
		return 0
	}

	bfs1(root.Right, nodes, true, 1)
	bfs1(root.Left, nodes, false, 1)
	maxDistance := 0
	startNode := nodes[start]
	for _, v := range nodes {
		if startNode.isLeft == v.isLeft {
			distance := int(math.Abs(float64(startNode.level - v.level)))
			if distance > maxDistance {
				maxDistance = distance
			}
		} else {
			distance := startNode.level + v.level
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}
	if startNode.level > maxDistance {
		return startNode.level
	}
	return maxDistance
}

func bfs1(root *TreeNode, nodes map[int]key, isLeft bool, level int) {
	if root == nil {
		return
	}
	nodes[root.Val] = key{
		level:  level,
		isLeft: isLeft,
	}
	bfs1(root.Left, nodes, isLeft, level+1)
	bfs1(root.Right, nodes, isLeft, level+1)
}
