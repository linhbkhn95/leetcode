package main

import "fmt"

type Node struct {
	Pre  *Node
	Next *Node
	Val  int
	Key  int
}

func main() {
	/**
	 * Your LRUCache object will be instantiated and called as such:
	 * obj := Constructor(capacity);
	 * param_1 := obj.Get(key);
	 * obj.Put(key,value);
	 */
	lRUCache := Constructor(3)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	lRUCache.Put(4, 4)
	fmt.Println(lRUCache.Get(4)) // return -1 (not found)

	fmt.Println(lRUCache.Get(3)) // return 1
	// LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(2)) // return -1 (not found)
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	lRUCache.Put(5, 5)
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(2)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return -1 (not found)
	fmt.Println(lRUCache.Get(4)) // return -1 (not found)

	fmt.Println(lRUCache.Get(5)) // return -1 (not found)

	// return 4
	// lRUCache := Constructor(1)
	// lRUCache.Put(2, 1)           // cache is {1=1}
	// fmt.Println(lRUCache.Get(2)) // return 4

}

type LRUCache struct {
	head       *Node
	tail       *Node
	cap        int
	currentCap int
	nodes      map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		nodes: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.nodes[key]
	if !ok {
		return -1
	}
	if this.currentCap > 1 {
		this.swapToHead(node)
	}
	return node.Val

}

func (this *LRUCache) Put(key int, value int) {
	nde, ok := this.nodes[key]
	if !ok {
		node := &Node{
			Val: value,
			Key: key,
		}

		if this.currentCap == this.cap {
			tail := this.tail
			head := this.head
			this.head = node
			this.head.Next = head
			if tail.Pre == nil {
				this.tail = node
			} else {
				head.Pre = node
				this.tail = tail.Pre
				this.tail.Next = nil
			}
			delete(this.nodes, tail.Key)

		} else {
			if this.currentCap == 0 {
				this.tail = node
				this.head = node
			} else {
				node.Next = this.head
				this.head.Pre = node
				this.head = node
			}
			this.currentCap++
		}
		this.nodes[key] = node

		return
	}
	nde.Val = value
	this.swapToHead(nde)
}

func (this *LRUCache) swapToHead(node *Node) {
	if this.currentCap == 1 {
		return
	}
	if node.Pre == nil {
		return
	}
	nextNode, preNode := node.Next, node.Pre
	if preNode != nil {
		preNode.Next = nextNode
	}
	if nextNode != nil {
		nextNode.Pre = preNode
	}
	if node.Next == nil {
		this.tail = preNode
	}
	node.Pre = nil
	node.Next = this.head
	this.head.Pre = node
	this.head = node
}
