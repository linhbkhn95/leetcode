package topkfrequentelements

import "sort"

type FreQuentNumber struct {
	Num       int
	Frequency int
}

func topKFrequent(nums []int, k int) []int {
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num]++
	}
	resultArr := &FreQuentArr{}
	for num, frequency := range mapping {
		(*resultArr) = append(*resultArr, &FreQuentNumber{
			Num:       num,
			Frequency: frequency,
		})
	}

	sort.Sort(resultArr)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = (*resultArr)[i].Num
	}
	return result
}

type FreQuentArr []*FreQuentNumber

// Len implements sort.Interface.
func (f FreQuentArr) Len() int {
	return len(f)
}

// Less implements sort.Interface.
func (f FreQuentArr) Less(i int, j int) bool {
	return f[i].Frequency > f[j].Frequency
}

// Swap implements sort.Interface.
func (f FreQuentArr) Swap(i int, j int) {
	f[i], f[j] = f[j], f[i]
}

var _ sort.Interface = (*FreQuentArr)(nil)
