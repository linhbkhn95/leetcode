package summaryrange

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := make([]string, 0)

	start, end := 0, 0
	for i := 1; i < len(nums); i++ {
		if i == end {
			continue
		}
		if nums[end]+1 == nums[i] {
			end++
			continue
		} else {
			if end-start == 0 {
				n := strconv.Itoa(nums[end])
				result = append(result, n)
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
			}
			start = i
			end = i
		}
	}

	if end-start == 0 {
		n := strconv.Itoa(nums[end])
		result = append(result, n)
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
	}
	return result
}
