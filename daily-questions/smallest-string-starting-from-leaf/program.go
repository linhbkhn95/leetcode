package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func main() {
	fmt.Println(string(byte(97)))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	var result string
	if root == nil {
		return ""
	}
	dfs(root, &result, "")
	return result
}

func dfs(root *TreeNode, result *string, currentString string) {
	currentString = string(byte(97+root.Val)) + currentString
	if root.Left == nil && root.Right == nil {
		if *result == "" {
			*result = currentString
		} else {
			*result = min(*result, currentString)
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, result, currentString)
	}
	if root.Right != nil {
		dfs(root.Right, result, currentString)
	}
}
