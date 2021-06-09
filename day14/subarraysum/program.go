package subarraysum

func subarraySum(nums []int, k int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		if nums[0] == k {
			return 1
		}
		return 0
	}
	sums := make([]int, len(nums))
	sums[0] = nums[0]

	for i := 1; i < l; i++ {
		sums[i] = nums[i] + sums[i-1]
	}
	result := 0
	for i := 0; i < l-1; i++ {
		if sums[i] == k {
			result++
		}
		for j := i + 1; j < l; j++ {
			if sums[j]-sums[i] == k {
				result++
			}
		}
	}
	if sums[l-1] == k {
		result++
	}
	return result
}

// type counter struct {
// 	value int
// }

// func subarraySumx(nums []int, k int) int {
// 	c := counter{}

// 	solution(nums, &c, false, len(nums), 0, 0, k)
// 	return c.value
// }

// func solution(nums []int, c *counter, hasItem bool, l, index, currentSum, k int) {

// 	if currentSum == k && hasItem {
// 		c.value += 1
// 	}
// 	if index == l-1 {
// 		if currentSum+nums[index] == k {
// 			c.value += 1
// 		}
// 		return
// 	}
// 	if hasItem {
// 		solution(nums, c, hasItem, l, index+1, currentSum+nums[index], k)
// 	} else {
// 		solution(nums, c, hasItem, l, index+1, currentSum, k)
// 		solution(nums, c, true, l, index+1, currentSum+nums[index], k)

// 	}

// }
