package gfg

import (
	"strconv"
)

func countAndSay(n int) string {

	return solution(n)
}

func solution(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}
	res := "11"
	for i := 2; i < n; i++ {
		res = say(res)
	}

	return res
}

// say use input "1211"
func say(n string) string {
	count := 0
	i := 0
	pre := n[0]
	result := ""
	for i < len(n) {
		if n[i] == pre {
			count++
		} else {
			result += strconv.Itoa(count) + string(pre)
			count = 1
			pre = n[i]
		}
		i++
	}
	result += strconv.Itoa(count) + string(pre)
	return result
}
