package sum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	result := make([][]int, 0)
	i, l := 0, len(nums)
	for i < l-1 {
		j := i + 1
		for j < l-2 {
			start, end := j+1, l-1
			for start < end {
				if (nums[start] + nums[end] + nums[i] + nums[j]) == target {
					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
					startNumber := nums[start]
					start += 1
					for start < l {
						if nums[start] != startNumber {
							break
						}
						start++
					}
					endNumber := nums[end]
					end -= 1
					for end >= 0 {
						if nums[end] != endNumber {
							break
						}
						end--
					}
				} else if (nums[start] + nums[end] + nums[i] + nums[j]) > target {
					end--
				} else {
					start++
				}
			}
			for j < l-1 {
				if nums[j] != nums[j+1] {
					break
				}
				j++
			}
		}
		for i < l-1 {
			if nums[i] != nums[i+1] {
				break
			}
			i++
		}

	}
	return result
}
