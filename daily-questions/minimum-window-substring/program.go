package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println('a')
}
func minWindow(s string, t string) string {
	countT := [58]int{}
	for i := range t {
		countT[t[i]-65]++
	}

	start, end, l := 0, 1, len(s)
	currentLen, minLen := 0, math.MaxInt
	windowCount := [58]int{}
	result := ""
	for start < l && countT[s[start]-65] == 0 {
		start++
	}
	if start < l {
		windowCount[s[start]-65]++
	} else {
		return ""
	}
	if len(t) == 1 && windowCount == countT {
		return t
	}
	end = start + 1
	duplicateCount := make(map[byte]int)
	for end < l {
		if countT[s[end]-65] != 0 {
			if countT[s[end]-65] > windowCount[s[end]-65] {
				windowCount[s[end]-65]++
			} else {
				duplicateCount[s[end]-65]++
			}
		}
		for windowCount == countT {
			currentLen = end - start + 1
			if currentLen < minLen {
				result = s[start : end+1]
				minLen = currentLen
			}
			if duplicateCount[s[start]-65] == 0 {
				windowCount[s[start]-65]--
			} else {
				duplicateCount[s[start]-65]--
			}
			start = start + 1
			for start < end && countT[s[start]-65] == 0 {
				start++
			}
		}
		end++
	}
	return result
}
