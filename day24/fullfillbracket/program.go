package fullfillbracket

import "fmt"

func Solution(s string) string {
	left, right := 0, 0
	result := s
	for i := 0; i < len(s); i++ {
		fmt.Println("ss", s[i] == '<')
		if s[i] == '<' {
			left++
			continue
		}
		if left > 0 {
			left--
		} else {
			right++
		}

	}
	fmt.Println("left, right", left, right)
	for j := 0; j < right; j++ {
		result = "<" + result
	}
	for j := 0; j < left; j++ {
		result += ">"
	}
	return result
}
