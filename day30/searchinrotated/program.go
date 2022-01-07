package program

import "fmt"

func search(nums []int, target int) int {
	return doSolution(nums, 0, len(nums)-1, target)
}

func doSolution(nums []int, start, end, target int) int {

	for start <= end {
		middleOffset := (start + end) / 2
		if nums[middleOffset] == target {
			return middleOffset
		} else if nums[middleOffset] > target {

			if nums[middleOffset] > nums[start] {
				index := binarySearch(nums, start, middleOffset-1, target)
				if index != -1 {
					return index
				}
			}
			index := doSolution(nums, middleOffset+1, end, target)
			if index != -1 {
				return index
			}
			return doSolution(nums, start, middleOffset-1, target)
		} else {

			if nums[middleOffset] < nums[end] {
				index := binarySearch(nums, middleOffset+1, end, target)
				if index != -1 {
					return index
				}
			}
			index := doSolution(nums, start, middleOffset-1, target)
			if index != -1 {
				return index
			}
			return doSolution(nums, middleOffset+1, end, target)
		}
	}
	return -1
}

func binarySearch(nums []int, start, end, target int) int {
	fmt.Println("binarySearch", start, end, target)
	for start <= end {
		middleOffset := (start + end) / 2
		if nums[middleOffset] == target {
			return middleOffset
		} else if nums[middleOffset] > target {
			end = middleOffset - 1
		} else {
			start = middleOffset + 1
		}
	}
	return -1
}
