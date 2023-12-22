package findduplicatenumber

func findDuplicate(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = true
		} else {
			return n
		}
	}
	return -1
}
