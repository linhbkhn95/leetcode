package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {
	node1 := &Node{7, nil, nil}
	node2 := &Node{13, nil, nil}
	node2.Random = node1
	node3 := &Node{11, nil, nil}

	node4 := &Node{10, nil, nil}
	node5 := &Node{1, nil, nil}
	node3.Random = node4
	node4.Random = node2
	node5.Random = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	result := copyRandomList(node1)
	fmt.Println("result", result)
}

func copyRandomList(head *Node) *Node {
	var t *Node
	t = head

	nodes := make(map[*Node]*Node)
	for t != nil {
		new := &Node{
			Val: t.Val,
		}
		nodes[t] = new
		t = t.Next
	}
	t = head
	for t != nil {
		nodes[t].Next = nodes[t.Next]
		nodes[t].Random = nodes[t.Random]
		t = t.Next
	}
	return nodes[head]
}
