package main

import "fmt"

type (
	Node struct {
		value int
		key   int
		pre   *Node
		next  *Node
	}

	LRUCache struct {
		data        map[int]*Node
		currentSize int
		size        int
		head        *Node
		last        *Node
	}
)

func Constructor(capacity int) LRUCache {
	return LRUCache{
		data: make(map[int]*Node, capacity),
		size: capacity,
	}
}

func (this *LRUCache) Get(key int) int {

	v, ok := this.data[key]
	if !ok {
		return -1
	}
	this.moveToHead(v)
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.data[key]
	if !ok {
		node := Node{
			value: value,
			key:   key,
		}
		this.data[key] = &node

		if this.head == nil {
			this.currentSize = 1
			this.head = &node
			this.last = &node
			return
		}
		if this.currentSize == this.size {
			last := this.last
			if last != nil {
				preLast := last.next
				if preLast != nil {
					preLast.pre = nil
				}
				this.last = preLast
			}
			this.insertToHead(&node)
			if last != nil {
				delete(this.data, last.key)
			}
			return
		}

		this.insertToHead(&node)
		this.currentSize += 1
		return
	}
	this.moveToHead(v)
	v.value = value
}

func (this *LRUCache) insertToHead(node *Node) {
	head := this.head
	head.next = node
	node.pre = head
	this.head = node
}

func (this *LRUCache) moveToHead(node *Node) {
	preV := node.pre
	nextV := node.next

	if node != this.head {
		if preV != nil {
			preV.next = nextV
		}
		if nextV != nil {
			nextV.pre = preV
		}
		if preV == nil && nextV != nil {
			this.last = nextV
			nextV.pre = nil
		}
		oldHead := this.head
		oldHead.next = node
		node.pre = oldHead
		node.next = nil
		this.head = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 1)           // cache is {1=1}
	lRUCache.Put(3, 2)           // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(3)) // return 1
	fmt.Println(lRUCache.Get(2)) // return 1

	lRUCache.Put(4, 3) // cache is {1=1, 2=2}

	fmt.Println(lRUCache.Get(2)) // return 1
	fmt.Println(lRUCache.Get(3)) // returns -1 (not found)
	fmt.Println(lRUCache.Get(4)) // return 1

	// lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	// lRUCache.Get(3)    // return 3
	// lRUCache.Get(4)    // return 4
}
