package main

import (
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= x
	})
	if index == -1 {
		if arr[0] > x {
			return arr[:k]
		}
		return arr[len(arr)-k:]
	}
	start := index - 1
	end := index
	count := k
	for start >= 0 && end < len(arr) && count != 0 {
		if abs(arr[start]-x) < abs(arr[end]-x) {
			start--
		} else if abs(arr[start]-x) > abs(arr[end]-x) {
			end++
		} else {
			start--
		}
		count--
	}

	if count != 0 {
		if start == -1 {
			for count != 0 {
				end++
				count--
			}
		} else {
			for count != 0 {
				start--
				count--
			}
		}
	}
	if start == -1 {
		return arr[:end]
	}
	if end == len(arr) {
		return arr[start+1:]
	}
	return arr[start+1 : end]
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
