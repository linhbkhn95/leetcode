package threesum

import "sort"

func threeSum(arr []int, k int) [][]int {
	l := len(arr)
	if l < 3 {
		return nil
	}

	sort.Ints(arr)

	var result [][]int
	for i := 0; i < l; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		j := i + 1
		h := l - 1
		for j < h {
			sum := arr[i] + arr[j] + arr[h]
			if sum == k {
				result = append(result, []int{arr[i], arr[j], arr[h]})
				q := j
				for q < l && arr[q] == arr[j] {
					q++
				}
				j = q
			} else if sum > k {
				h--
			} else {
				j++
			}
		}
	}

	return result
}
