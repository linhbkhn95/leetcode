package convertanarrayintoa2darraywithconditions

import "sort"

func findMatrix(nums []int) [][]int {
	freq := make(map[int]int, 0)
	result := make([][]int, 0)
	for _, num := range nums {
		if freq[num] >= len(result) {
			result = append(result, []int{})
		}
		result[freq[num]] = append(result[freq[num]], num)

		freq[num]++
	}
	return result
}

func findMatrix1(nums []int) [][]int {
	result := make([][]int, 0)
	index := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	lastN := nums[0]
	result = append(result, []int{nums[0]})
	for i := 1; i < len(nums); i++ {
		if lastN == nums[i] {
			if len(result)-1 > index {
				index += 1
				result[index] = append(result[index], nums[i])
			} else {
				result = append(result, []int{nums[i]})
				index++
			}
		} else {
			index = 0
			result[index] = append(result[index], nums[i])
			lastN = nums[i]

		}

	}
	return result
}
