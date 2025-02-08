package main

import "fmt"

type NumberContainers struct {
	numbers map[int]*Node
	indexes map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		numbers: make(map[int]*Node),
		indexes: make(map[int]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	oldValue, ok := this.indexes[index]
	if ok && oldValue == number {
		return
	}
	if ok && oldValue != number {
		btree, _ := this.numbers[oldValue]
		this.numbers[oldValue] = btree.Delete(index)
	}
	this.numbers[number] = this.numbers[number].Insert(index)
	this.indexes[index] = number
}

func (this *NumberContainers) Find(number int) int {
	btree, ok := this.numbers[number]
	if !ok {
		return -1
	}
	if btree == nil {
		return -1
	}
	min := btree.FindMin()
	if min == nil {
		return -1
	}
	return min.Value
}

// Node represents a single node in the BST
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert inserts a new value into the BST
func (n *Node) Insert(value int) *Node {
	if n == nil {
		return &Node{Value: value}
	}
	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
	}
	return n
}

// Search checks if a value exists in the BST
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// FindMin finds the node with the smallest value in the BST
func (n *Node) FindMin() *Node {
	if n.Left == nil {
		return n
	}
	return n.Left.FindMin()
}

// Delete removes a value from the BST
func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}
	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		// Case 1: No child
		if n.Left == nil && n.Right == nil {
			return nil
		}
		// Case 2: One child
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}
		// Case 3: Two children
		minRight := n.Right.FindMin()
		n.Value = minRight.Value
		n.Right = n.Right.Delete(minRight.Value)
	}
	return n
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */

func main() {
	nc := Constructor()
	nc.Change(1, 10)         // Your container at index 2 will be filled with number 10.
	fmt.Println(nc.Find(10)) // There is no index that is filled with number 10. Therefore, we return -1.
	nc.Change(1, 20)         // Your container at index 3 will be filled with number 10.
	fmt.Println(nc.Find(10)) // There is no index that is filled with number 10. Therefore, we return -1.
	fmt.Println(nc.Find(20)) // There is no index that is filled with number 10. Therefore, we return -1.
	fmt.Println(nc.Find(30)) // There is no index that is filled with number 10. Therefore, we return -1.

}
