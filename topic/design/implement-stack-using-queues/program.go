package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Top())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Empty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())

}

type MyStack struct {
	q1, q2 *list.List
}

func Constructor() MyStack {
	return MyStack{
		list.New(), list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.q1.PushBack(x)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	item, q := this.Peak()
	q.Remove(item)
	return item.Value.(int)
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	item, q := this.Peak()
	if this.q1 == q {
		q.Remove(item)
		this.q2.PushBack(item.Value)
	} else {
		q.Remove(item)
		this.q1.PushBack(item.Value)
	}
	return item.Value.(int)

}

func (this *MyStack) Peak() (*list.Element, *list.List) {
	for this.q1.Len() > 1 {
		this.q2.PushBack(this.q1.Front().Value)
		this.q1.Remove(this.q1.Front())
	}
	if this.q1.Len() == 1 {
		item := this.q1.Front()
		return item, this.q1
	}
	for this.q2.Len() > 1 {
		this.q1.PushBack(this.q2.Front().Value)
		this.q2.Remove(this.q2.Front())
	}
	if this.q2.Len() == 1 {

		item := this.q2.Front()

		return item, this.q2
	}
	return nil, nil
}

func (this *MyStack) Empty() bool {
	return this.q1.Len() == 0 && this.q2.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
