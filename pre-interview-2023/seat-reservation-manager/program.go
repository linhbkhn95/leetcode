package seatreservationmanager

import "container/heap"

type SeatManager struct {
	piq PriorityQueue
}

func Constructor(n int) SeatManager {
	piq := make(PriorityQueue, 0)
	for i := 1; i <= n; i++ {
		heap.Push(&piq, i)
	}
	return SeatManager{
		piq,
	}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(&this.piq).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.piq, seatNumber)

}

type PriorityQueue []int

func (piq PriorityQueue) Len() int {
	return len(piq)
}
func (piq PriorityQueue) Less(i, j int) bool {

	return piq[i] < piq[j]
}
func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
}
func (piq *PriorityQueue) Push(x interface{}) {
	item := x.(int)
	*piq = append(*piq, item)
}
func (piq *PriorityQueue) Pop() interface{} {
	old := *piq
	n := len(old)
	item := old[n-1]
	*piq = old[0 : n-1]
	return item
}

func (piq PriorityQueue) Top() interface{} {
	return piq[0]
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

//------------------------------------------------------------------------------

type BST struct {
	size int
	root *Node
}

//------------------------------------------------------------------------------

func (tree *BST) insert(data int) {

	tree.root = tree.add(data, tree.root)
}

//------------------------------------------------------------------------------

func (tree *BST) add(data int, node *Node) *Node {

	if node == nil {

		tree.size++

		return &Node{data: data}
	}

	if data < node.data {

		node.left = tree.add(data, node.left)

	} else if data > node.data {

		node.right = tree.add(data, node.right)
	}

	return node
}

//------------------------------------------------------------------------------

func (tree *BST) remove(data int) {

	tree.root = tree.delete(data, tree.root)
}

//------------------------------------------------------------------------------

func (tree *BST) delete(data int, node *Node) *Node {

	if node == nil {

		return nil
	}

	if data < node.data {

		node.left = tree.delete(data, node.left)

	} else if data > node.data {

		node.right = tree.delete(data, node.right)

	} else {

		if node.left != nil && node.right != nil {

			min := tree.min(node.right)
			node.data = min.data
			node.right = tree.delete(min.data, node.right)

		} else {

			if node.left == nil && node.right == nil {

				node = nil

			} else if node.left == nil {

				node = node.right

			} else {

				node = node.left
			}

			tree.size--
		}
	}

	return node
}

//------------------------------------------------------------------------------

func (tree *BST) min(node *Node) *Node {

	for node.left != nil {

		node = node.left
	}

	return node
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
