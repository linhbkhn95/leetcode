package line

import (
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Top() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func Solution(S string) int {
	// write your code in Go 1.4
	arr := strings.Split(S, " ")
	stack := &Stack{}
	for _, ac := range arr {
		if ac == "POP" {
			stack.Pop()
		} else if ac == "DUP" {
			num, ok := stack.Top()
			if !ok {
				return -1
			}
			stack.Push(num)
		} else if ac == "-" || ac == "+" {
			num1, ok := stack.Pop()
			if !ok {
				return -1
			}
			num2, ok := stack.Pop()
			if !ok {
				return -1
			}
			if ac == "+" {
				stack.Push(num1 + num2)
			} else {
				stack.Push(num1 - num2)
			}
		} else {
			num, err := strconv.Atoi(ac)
			if err != nil {
				return -1
			}
			stack.Push(num)
		}
	}
	num, ok := stack.Pop()
	if !ok {
		return -1
	}
	return num
}

func Solution1(S string) int {
	// write your code in Go 1.4
	logs := strings.Split(S, "\n")
	mapping := make(map[string]int)
	for _, log := range logs {
		data := strings.Split(log, ",")

		duration := data[0]
		phone := data[1]
		sum, ok := mapping[phone]
		sencondDuration := convertToSecond(duration)
		if ok {
			mapping[phone] = sum + sencondDuration
		} else {
			mapping[phone] = sencondDuration
		}
	}

	log := logs[0]
	data := strings.Split(log, ",")
	minPhone := data[1]
	min := mapping[minPhone]
	for phone, sum := range mapping {
		if sum < min {
			minPhone = phone
			min = sum
		}
	}
	delete(mapping, minPhone)
	result := 0
	for _, sum := range mapping {
		if sum <= 300 {
			result += sum * 3
		}
		min := sum / 60
		return min * 150
	}
	return result
}

func convertToSecond(duration string) int {
	result := 0
	data := strings.Split(duration, ":")
	second, _ := strconv.Atoi(data[2])
	result += second
	min, _ := strconv.Atoi(data[1])
	result += 60 * min
	hour, _ := strconv.Atoi(data[0])
	result += 3600 * hour
	return result
}
