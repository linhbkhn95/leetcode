package parentheses

// func longestValidParentheses(s string) int {
// 	str := ""
// 	max := 0
// 	currentCount := 0
// 	preCount := 0
// 	for i, c := range s {
// 		if len(str) == 0 {
// 			currentCount += preCount
// 			preCount = currentCount
// 			if max < currentCount {
// 				max = currentCount
// 			}
// 			currentCount = 0
// 		}
// 		if c == '(' {
// 			if currentCount > 0 {
// 				preCount = currentCount
// 				currentCount = 0
// 			}

// 			str += string(c)
// 		} else {
// 			if len(str) == 0 {
// 				currentCount += preCount
// 				if max < currentCount {
// 					max = currentCount
// 				}
// 				currentCount = 0
// 				preCount = 0
// 				continue
// 			}
// 			str = str[:len(str)-1]
// 			currentCount += 2

// 			if i == len(s)-1 {
// 				if len(str) == 0 {
// 					currentCount += preCount
// 				}
// 			}
// 		}
// 	}
// 	if currentCount > max {
// 		return currentCount
// 	}
// 	return max
// }

type Node struct {
	Val  int
	Next *Node
}

type Stack struct {
	Head *Node
}

func (stack *Stack) Push(c int) {
	node := &Node{
		Val: c,
	}
	head := stack.Head
	node.Next = head
	stack.Head = node
}

func (stack *Stack) Pop() int {
	if stack.IsEmpty() {
		return '0'
	}
	c := stack.Head.Val
	stack.Head = stack.Head.Next
	return c
}

func (stack *Stack) IsEmpty() bool {
	return stack.Head == nil
}

type target struct {
	Start int
	End   int
	Next  *target
}

type TargetStack struct {
	Head *target
}

func (stack *TargetStack) Push(target *target) {
	head := stack.Head
	target.Next = head
	stack.Head = target
}

func (stack *TargetStack) Pop() *target {
	if stack.IsEmpty() {
		return nil
	}
	c := stack.Head
	stack.Head = stack.Head.Next
	return c
}

func (stack *TargetStack) IsEmpty() bool {
	return stack.Head == nil
}

func longestValidParentheses1(s string) int {
	result := &TargetStack{}
	stack := &Stack{}
	max := 0
	for i, c := range s {
		if c == '(' {
			stack.Push(i)
		} else {
			if stack.IsEmpty() {
				continue
			}
			index := stack.Pop()
			if result.IsEmpty() {

				result.Push(&target{
					Start: index,
					End:   i,
				})
				max = maxInt(max, i-index+1)

			} else {
				t := result.Head
				for t != nil && index < t.Start {
					result.Pop()
					t = result.Head
				}
				if t != nil && index-t.End == 1 {
					t := result.Pop()
					if t != nil {
						result.Push(&target{
							Start: t.Start,
							End:   i,
						})
						max = maxInt(max, i-t.Start+1)
					}

				} else {
					max = maxInt(max, i-index+1)
					result.Push(&target{
						Start: index,
						End:   i,
					})
				}

			}
		}
	}
	return max
}

func maxInt(first, second int) int {
	if first > second {
		return first
	}
	return second
}
