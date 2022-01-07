package parentheses

func isValid(s string) bool {
	stack := Stack{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
		} else {
			if stack.IsEmpty() {
				return false
			}
			pre := stack.Pop()
			if pre == '(' && c == ')' || pre == '[' && c == ']' || pre == '{' && c == '}' {
				continue
			}
			return false
		}
	}
	return stack.IsEmpty()
}

type Node struct {
	Val  rune
	Next *Node
}

type Stack struct {
	Head *Node
}

func (stack *Stack) Push(c rune) {
	node := &Node{
		Val: c,
	}
	head := stack.Head
	node.Next = head
	stack.Head = node
}

func (stack *Stack) Pop() rune {
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

func isValid2(s string) bool {
	str := ""
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			str += string(c)
		} else {

			if len(str) == 0 {
				return false
			}
			pre := str[len(str)-1]
			if pre == '(' && c == ')' || pre == '[' && c == ']' || pre == '{' && c == '}' {
				str = str[:len(str)-1]
				continue
			}
			return false
		}
	}
	return len(str) == 0
}
