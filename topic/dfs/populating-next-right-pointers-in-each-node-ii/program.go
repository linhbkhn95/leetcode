package populatingnextrightpointersineachnodeii

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	bag := make(map[int][]*Node, 0)
	traversal(root, bag, 0)
	for _, nodes := range bag {
		l := len(nodes)
		if l == 1 {
			continue
		}
		for i := 1; i < l; i++ {
			nodes[i-1].Next = nodes[i]
		}
	}
	return root
}

func traversal(root *Node, bag map[int][]*Node, level int) {
	if root == nil {
		return
	}
	bag[level] = append(bag[level], root)
	traversal(root.Left, bag, level+1)
	traversal(root.Right, bag, level+1)
}
