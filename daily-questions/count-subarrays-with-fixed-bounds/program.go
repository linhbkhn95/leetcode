package count_subarrays_with_fixed_bounds

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64
	left := 0
	pmin := -1
	pmax := -1

	for right, num := range nums {
		if num < minK || num > maxK {
			left = right + 1
			pmin = -1
			pmax = -1
		} else {
			if num == minK {
				pmin = right
			}
			if num == maxK {
				pmax = right
			}
			res += int64(max(0, min(pmin, pmax)-left+1))
		}
	}

	return res
}

// 3 4 5 5 3
