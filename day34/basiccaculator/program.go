package program

import (
	"fmt"
	"strconv"
)

var operators = map[string]interface{}{
	"-": nil,
	"+": nil,
	"*": nil,
	"/": nil,
}

func calculate(s string) int {
	l := len(s)
	result := []string{}
	i := 0
	for i < len(s) {
		c := s[i]
		if c == ' ' {
			i++
			continue
		}
		if !isOperator(string(c)) {
			j := extractNextNumber(s, l, i)
			result = append(result, s[i:j])
			i = j
		} else {
			if c == '*' || c == '/' {

				pre, _ := strconv.Atoi(string(result[len(result)-1]))
				result = result[:len(result)-1]
				i += 1
				nextC := s[i]
				for nextC == ' ' {
					i = i + 1
					nextC = s[i]
				}
				j := extractNextNumber(s, l, i)
				next, _ := strconv.Atoi(s[i:j])
				i = j
				new := processOperator(pre, next, string(c))
				result = append(result, fmt.Sprintf("%d", new))
			} else {
				result = append(result, string(c))
				i++
			}
		}
	}
	total := 0
	k := 0
	for k < len(result) {
		c := result[k]
		if !isOperator(c) {
			pre, _ := strconv.Atoi(string(c))
			total += pre
			k++
		} else {
			next := result[k+1]
			n, _ := strconv.Atoi(next)
			total = processOperator(total, n, c)
			k += 2
		}
	}
	return total
}

func extractNextNumber(s string, l, i int) int {
	j := i + 1

	for j < l && !isOperator(string(s[j])) && s[j] != ' ' {
		j++
	}
	return j
}

func isOperator(c string) bool {
	_, ok := operators[c]
	return ok
}

func processOperator(first, second int, c string) int {
	switch c {
	case "*":
		{
			return first * second
		}
	case "/":
		{
			return first / second
		}
	case "+":
		{
			return first + second
		}
	default:
		return first - second
	}
}
