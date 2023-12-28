package longestcontinuoussubarraywithabsolutedifflessthanorequaltolimit

import (
	"container/list"
)

func longestSubarray(nums []int, limit int) int {
	if len(nums) < 2 {
		return 1
	}
	min := list.New()
	max := list.New()

	min.PushBack(0)
	max.PushBack(0)

	start, end := 0, 1
	l := len(nums)
	result := 1
	for end < l {
		current := nums[end]
		for min.Len() > 0 && nums[min.Back().Value.(int)] >= current {
			min.Remove(min.Back())
		}
		for max.Len() > 0 && nums[max.Back().Value.(int)] <= current {
			max.Remove(max.Back())
		}

		min.PushBack(end)
		max.PushBack(end)

		for (nums[max.Front().Value.(int)] - nums[min.Front().Value.(int)]) > limit {
			start++
			for max.Len() > 0 && max.Front().Value.(int) < start {
				max.Remove(max.Front())
			}
			for min.Len() > 0 && min.Front().Value.(int) < start {
				min.Remove(min.Front())
			}
		}
		result = maxInt(result, end-start+1)
		end++
	}
	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
