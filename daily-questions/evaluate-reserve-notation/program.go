package evaluatereservenotation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := NewStack()
	var result int
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token == "+" || token == "-" || token == "*" || token == "/" {
			num1 := stack.Pop()
			num2 := stack.Pop()
			switch token {
			case "+":
				{
					result = num1 + num2
				}
			case "-":
				{
					result = num2 - num1
				}
			case "/":
				{
					result = num2 / num1
				}
			case "*":
				{
					result = num1 * num2
				}
			}
			stack.Push(result)
		} else {
			num, _ := strconv.ParseInt(tokens[i], 10, 64)
			stack.Push(int(num))
		}
	}
	return stack.Pop()
}

type stack struct {
	arr []int
}

func NewStack() *stack {
	return &stack{
		arr: make([]int, 0),
	}
}

func (s *stack) Push(number int) {
	s.arr = append(s.arr, number)
}

func (s *stack) Pop() int {
	num := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return num
}

func (s stack) IsEmpty() bool {
	return len(s.arr) == 0
}
