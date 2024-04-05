package main

import (
	"fmt"
)

func main() {
	fmt.Println('A', 'a')
}

func makeGood(s string) string {
	start := 0
	result := make([]byte, 0, len(s))
	for start < len(s) {
		if len(result) == 0 {
			result = append(result, s[start])
		} else {
			lastC := result[len(result)-1]
			val := abs(int(s[start]) - int(lastC))
			if val == 32 {
				result = result[:len(result)-1]
			} else {
				result = append(result, s[start])
			}
		}
		start++
	}
	return string(result)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

/*
Example 1:

Input: s = "leEeetcode"
Output: "leetcode"
Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
*/
