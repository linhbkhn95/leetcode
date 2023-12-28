package main

import "fmt"

func main() {
	fmt.Println(3 & 1)
}

func longestNiceSubarray(nums []int) int {
	maxLen := 0
	l := len(nums)
	start, end := 0, 1
	currentLen := 1
	for start < l-1 && end < l {
		ok, index := isA(nums, start, end)
		if !ok {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			start = index
			end = start + 1
		} else {
			currentLen = end - start + 1
			end++
		}
	}
	if currentLen > maxLen {
		maxLen = currentLen
	}

	return maxLen
}

func isA(nums []int, start, end int) (bool, int) {
	for i := end - 1; i >= start; i-- {
		if nums[i]&nums[end] != 0 {
			return false, i + 1
		}
	}
	return true, -1
}
