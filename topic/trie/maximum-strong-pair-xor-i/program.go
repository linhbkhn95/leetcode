package maximumstrongpairxori

import "math"

func maximumStrongPairXor(nums []int) int {
	maxV := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i])-float64(nums[j]))) <= min(nums[i], nums[j]) {
				maxV = max(maxV, nums[i]^nums[j])
			}
		}
	}
	return maxV
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
