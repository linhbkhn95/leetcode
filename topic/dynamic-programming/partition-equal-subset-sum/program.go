package partitionequalsubsetsum

import (
	"fmt"
	"sort"
)

func canPartition(nums []int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}
	if s%2 == 1 {
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	wrongWays := make(map[string]struct{})
	return tryPartition(nums, len(nums), 0, s/2, wrongWays)

}

func tryPartition(nums []int, n, index, target int, wrongWays map[string]struct{}) bool {
	if _, ok := wrongWays[fmt.Sprintf("%d-%d", index, target)]; ok {
		return false
	}
	if target < 0 {
		return false
	}
	if index == n {
		return false
	}
	if target == 0 {
		return true
	}

	ok := tryPartition(nums, n, index+1, target-nums[index], wrongWays)
	if ok {
		return true
	}
	wrongWays[fmt.Sprintf("%d-%d", index, target)] = struct{}{}
	return tryPartition(nums, n, index+1, target, wrongWays)
}
