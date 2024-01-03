package main

import (
	"fmt"
)

func main() {
	fmt.Println('a', 'z', 'A', 'Z', ' ', ',', '0', '1', '9', ':', '_')
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	var abs int
	for i < j {
		fmt.Println("s[i]", s[i])
		fmt.Println("s[j]", s[j])

		if !isEligibleChar(s[i]) {
			i++
			continue
		}
		if !isEligibleChar(s[j]) {
			j--
			continue
		}

		if s[i] == s[j] {
			i++
			j--
			continue
		}

		if s[i] < 64 {
			return false
		}
		if s[i] > s[j] {
			abs = int(s[i]) - int(s[j])
		} else {
			abs = int(s[j]) - int(s[i])
		}
		if abs != 32 {
			return false
		}
		i++
		j--
	}
	return true
}

func isEligibleChar(c byte) bool {
	return 64 < c && c < 91 || 96 < c && c < 123 || 47 < c && c < 58
}

// ' '
