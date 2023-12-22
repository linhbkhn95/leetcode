package main

type (
	MinStack struct {
		arr  []*Node
		head *Node
		len  int
	}

	Node struct {
		val  int
		next *Node
		pre  *Node
	}
)

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	node := Node{
		val: val,
	}
	if this.head == nil {
		this.head = &node
	} else {
		head := this.head
		pre := this.head.pre
		for head != nil && node.val > head.val {
			pre = head
			head = head.next
		}
		if pre == nil {
			h := this.head
			node.next = h
			h.pre = &node
			this.head = &node
		}
		if head == nil {
			pre.next = &node
			node.pre = pre
		}
	}
	this.len += 1
	this.arr = append(this.arr, &node)
}

func (this *MinStack) Pop() {
	node := this.arr[this.len-1]
	nextNode := node.next
	preNode := node.pre

	if node == this.head {
		if nextNode != nil {
			nextNode.pre = nil
		}
		this.head = nextNode
	} else {
		if nextNode != nil {
			nextNode.pre = preNode
		}
		if preNode != nil {
			preNode.next = nextNode
		}
	}
	this.len -= 1
	this.arr = this.arr[0:this.len]
}

func (this *MinStack) Top() int {
	return this.arr[this.len-1].val
}

func (this *MinStack) GetMin() int {
	if this.head != nil {
		return this.head.val
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin() // return -3
	minStack.Pop()
	minStack.Top()    // return 0
	minStack.GetMin() // return -2
}
