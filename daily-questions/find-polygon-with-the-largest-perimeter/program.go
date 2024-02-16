package findpolygonwiththelargestperimeter

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	var sum int64
	sum = int64(nums[0]) + int64(nums[1])
	var result int64
	result = -1
	for i := 2; i < len(nums); i++ {
		if int64(nums[i]) < sum {
			result = max(result, sum+int64(nums[i]))
		}
		sum += int64(nums[i])
	}
	return result
}
