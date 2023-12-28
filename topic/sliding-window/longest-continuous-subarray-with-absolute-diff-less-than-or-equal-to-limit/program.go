package longestcontinuoussubarraywithabsolutedifflessthanorequaltolimit

import (
	"container/list"
	"math"
)

func longestSubarray2(nums []int, limit int) int {
	res := 1
	n := len(nums)
	mn := list.New()
	mx := list.New()
	mn.PushBack(0)
	mx.PushBack(0)
	l := 0
	r := 1
	for r < n {
		cur := nums[r]
		for mn.Len() > 0 && nums[mn.Back().Value.(int)] >= cur {
			mn.Remove(mn.Back())
		}
		for mx.Len() > 0 && nums[mx.Back().Value.(int)] <= cur {
			mx.Remove(mx.Back())
		}
		mn.PushBack(r)
		mx.PushBack(r)
		for nums[mx.Front().Value.(int)]-nums[mn.Front().Value.(int)] > limit {
			l++
			for mx.Front().Value.(int) < l {
				mx.Remove(mx.Front())
			}
			for mn.Front().Value.(int) < l {
				mn.Remove(mn.Front())
			}
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func longestSubarray1(nums []int, limit int) int {

	l := len(nums)
	if l < 2 {
		return 1
	}
	maxLen := 0

	start, end := 0, 1
	currentLen := 1

	for start < l-1 && end < l {
		index, ok := checkCondition(nums, start, end, limit)
		if !ok {
			start = index
			end = index + 1
			maxLen = max(maxLen, currentLen)
			currentLen = 1
			continue
		}
		currentLen += 1
		end++
	}

	return max(maxLen, currentLen)
}

func checkCondition(nums []int, start, end int, limit int) (int, bool) {
	for i := end - 1; i >= start; i-- {
		if int(math.Abs(float64(nums[i]-nums[end]))) > limit {
			return i + 1, false
		}
	}
	return 0, true
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Extract(arr []int) (int, int) {
	min, max := arr[0], arr[0]
	for _, num := range arr {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}
	return min, max

}

// [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,
// [10,1,2,4,7,2]
// -> [1,2,2,4,7,10]
