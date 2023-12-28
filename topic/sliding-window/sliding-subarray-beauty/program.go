package slidingsubarraybeauty

import "sort"

func getSubarrayBeauty(nums []int, k int, x int) []int {

	l := len(nums)
	result := []int{}
	sortArr := make([]int, k)
	lastIndex := -1
	sortedArr := &SortedArr{}
	for i := 0; i < k; i++ {
		sort.
	}

	for start := 0; start < l-k+1; start++ {
		if lastIndex == -1 {
			// if nums[start+k-1] > nums[k]
		}
		for i := 0; i < k; i++ {
			sortArr[i] = nums[start+i]
		}
		sort.Slice(sortArr, func(i, j int) bool {
			return sortArr[i] < sortArr[j]
		})
		if sortArr[x-1] >= 0 {
			result = append(result, 0)
		} else {
			result = append(result, sortArr[x-1])
		}
	}
	return result
}

type Element struct {
	Value int
	Index int
}

type SortedArr []*Element

// Len implements sort.Interface.
func (s SortedArr) Len() int {
	return len(s)
}

// Less implements sort.Interface.
func (s SortedArr) Less(i int, j int) bool {
	return s[i].Value < s[j].Value
}

// Swap implements sort.Interface.
func (s SortedArr) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = (*SortedArr)(nil)
