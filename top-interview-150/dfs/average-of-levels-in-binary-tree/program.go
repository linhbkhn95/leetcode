package main

import "math"

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
type info struct {
	count int
	sum   int
}

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	levelInfo := make(map[int]*info, 0)
	dfs(root, 0, &result, levelInfo)
	for i := range result {
		info := levelInfo[i]
		result[i] = roundFloat(float64(info.sum)/float64(info.count), 5)
	}
	return result
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func dfs(root *TreeNode, level int, result *[]float64, levelInfo map[int]*info) {
	if root == nil {
		return
	}

	if level >= len(*result) {
		*result = append(*result, 0)
	}
	i, ok := levelInfo[level]
	if !ok {
		levelInfo[level] = &info{
			count: 1,
			sum:   root.Val,
		}
	} else {
		i.count++
		i.sum += root.Val
	}
	dfs(root.Left, level+1, result, levelInfo)
	dfs(root.Right, level+1, result, levelInfo)
}
