package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	dp(nums, len(nums), 0, []int{}, maxArr)
}

func dp(nums []int, l, index int, bag []int, maxArr *[]int) {
	if index == l {
		if len(*maxArr) < len(bag) {
			*maxArr = bag
		}
		return
	}
	if nums[index]%nums[index+1] == 0 {
		dp(nums, l, index+1, append(bag, nums[index+1]), maxArr)
	} else {
		dp(nums, l, index+1, []int{nums[index+1]}, maxArr)
	}
}

// 3 2 1
//
