package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := len(nums)
	if l%3 != 0 {
		return nil
	}
	result := make([][]int, 0)
	for i := 0; i < l-2; i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		result = append(result, nums[i:i+3])
	}
	return result
}
