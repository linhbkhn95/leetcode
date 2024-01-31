package main

import (
	"container/list"
	"fmt"
)

func main() {
	myQueue := Constructor()
	myQueue.Push(1)              // queue is: [1]
	myQueue.Push(2)              // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Println(myQueue.Peek())  // return 1
	fmt.Println(myQueue.Pop())   // return 1, queue is [2]
	fmt.Println(myQueue.Empty()) // return false
}

type MyQueue struct {
	stack1 *list.List
	stack2 *list.List
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1.PushBack(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	x := this.Peek()
	this.stack2.Remove(this.stack2.Back())
	return x
}

func (this *MyQueue) Peek() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	x := this.stack2.Back()
	return x.Value.(int)
}

func (this *MyQueue) Empty() bool {
	return this.stack1.Len() == 0 && this.stack2.Len() == 0
}
