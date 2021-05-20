package generateparent

type Bag struct {
	items []string
}

func generateParenthesis(n int) []string {
	// use to coller result
	bag := &Bag{
		items: nil,
	}
	// use to calc result
	var stack Stack
	findParenthesis(bag, stack, n, n, "(", "(", 1, 0)
	return bag.items
}

// using tree to find result.
func findParenthesis(bag *Bag, stack Stack, NumofParenthesisNeedcompleted int, n int, current string, pt string, left, right int) {
	// cut branches that can not generate result.
	if left-1 > n || right-1 > n {
		return
	}
	if pt == "(" {
		if NumofParenthesisNeedcompleted > 0 {
			stack.Push(pt)
			findParenthesis(bag, stack, NumofParenthesisNeedcompleted, n, current+")", ")", left, right+1)
			findParenthesis(bag, stack, NumofParenthesisNeedcompleted, n, current+"(", "(", left+1, right)
		}
		return
	}
	char, ok := stack.Pop()
	if ok && char == "(" {
		NumofParenthesisNeedcompleted -= 1
		if NumofParenthesisNeedcompleted == 0 {
			if stack.IsEmpty() {
				bag.items = append(bag.items, current)
			}

		} else {
			findParenthesis(bag, stack, NumofParenthesisNeedcompleted, n, current+")", ")", left, right+1)
			findParenthesis(bag, stack, NumofParenthesisNeedcompleted, n, current+"(", "(", left+1, right)

		}
	}
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
