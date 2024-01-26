package setmismatch

func findErrorNums(nums []int) []int {
	sets := make(map[int]int)
	for i := 1; i <= len(nums); i++ {
		sets[i]++
	}

	for _, n := range nums {
		sets[n]--
		if sets[n] == 0 {
			delete(sets, n)
		}
	}
	result := make([]int, 2)
	for k, v := range sets {
		if v == -1 {
			result[0] = k
		} else if v == 1 {
			result[1] = k
		}
	}
	return result
}
