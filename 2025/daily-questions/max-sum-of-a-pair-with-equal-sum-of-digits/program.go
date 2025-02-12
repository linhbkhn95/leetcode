package maxsumofapairwithequalsumofdigits

import "strconv"

func maximumSum(nums []int) int {
	top2Nums := make(map[int][2]int, 0)
	result := -1
	for _, num := range nums {
		sum := sumOfDigits(num)
		nums, ok := top2Nums[sum]
		if !ok {
			top2Nums[sum] = [2]int{num, 0}
			continue
		}
		if nums[1] == 0 {
			nums[1] = num
		} else {
			if num <= nums[1] {
				continue
			}
			nums[1] = num
		}
		if num > nums[0] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		top2Nums[sum] = nums
		result = max(nums[0]+nums[1], result)

	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumOfDigits(num int) int {
	numStr := strconv.Itoa(num)
	result := 0
	for i := range numStr {
		num, _ := strconv.Atoi(string(numStr[i]))
		result += num
	}
	return result
}
