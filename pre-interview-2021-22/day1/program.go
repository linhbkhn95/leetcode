package day1

import (
	"math"
)

func Permutations(array []int) [][]int {
	return nil
}

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	return findThreeSumClosest(nums, 0, 3, 0, nums[0]+nums[1]+nums[2], target)
}

func findThreeSumClosest(nums []int, offset, k, currentSum, curentResult, target int) int {

	if k == 0 {
		if math.Abs(float64(curentResult)-float64(target)) > math.Abs(float64(currentSum)-float64(target)) {
			return currentSum
		}
		return curentResult
	} else {
		if offset == len(nums)-1 {
			if k == 1 {
				return currentSum + nums[offset]

			}
			return curentResult
		}
		case1 := findThreeSumClosest(nums, offset+1, k-1, currentSum+nums[offset], curentResult, target)
		case2 := findThreeSumClosest(nums, offset+1, k, currentSum, curentResult, target)

		if math.Abs(float64(case1)-float64(target)) > math.Abs(float64(case2)-float64(target)) {
			return case2
		}
		return case1
	}
}
