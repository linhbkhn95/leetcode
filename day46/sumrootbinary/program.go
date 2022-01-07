package sumrootbinary

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	bag := []int{root.Val}
	var sum int
	move(root, &sum, bag)
	return sum
}

func move(root *TreeNode, sum *int, bag []int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*sum += covertToNumber(bag)
		return
	}
	if root.Left != nil {
		bag = append(bag, root.Left.Val)
		move(root.Left, sum, bag)
		bag = bag[:len(bag)-1]
	}
	if root.Right != nil {
		bag = append(bag, root.Right.Val)
		move(root.Right, sum, bag)
		bag = bag[:len(bag)-1]
	}

}

func covertToNumber(bag []int) int {
	num := 0
	l := len(bag)
	for i := l - 1; i > -1; i-- {
		if bag[i] != 0 {
			num += int(math.Pow(2, float64(l-1-i)))
		}
	}
	return num
}
