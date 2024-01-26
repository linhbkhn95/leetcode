package main

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

func pseudoPalindromicPaths(root *TreeNode) int {
	result := 0
	dfs(root, make(map[int]int), &result)
	return result
}

func dfs(root *TreeNode, nodes map[int]int, count *int) {
	if root == nil {
		return
	}
	nodes[root.Val]++
	if root.Left == nil && root.Right == nil {
		if isPseudoPalindromic(nodes) {
			*count += 1
		}
	}
	dfs(root.Left, nodes, count)
	dfs(root.Right, nodes, count)

	nodes[root.Val]--

}

func isPseudoPalindromic(nodes map[int]int) bool {
	hasOdd := false
	for _, count := range nodes {
		if hasOdd && count%2 == 1 {
			return false
		}
		if count%2 == 1 {
			hasOdd = true
		}
	}
	return true
}
