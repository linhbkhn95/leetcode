package program

type Node struct {
	Next  *Node
	Left  *Node
	Right *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		solution(root.Left, root.Right)
	}
	return root
}

func solution(left, right *Node) {
	if left == nil {
		return
	}
	left.Next = right
	t, r := left, right
	for t.Left != nil {
		t.Right.Next = r.Left
		t = t.Right
		r = r.Left
	}
	solution(left.Left, left.Right)
	solution(right.Left, right.Right)
}
