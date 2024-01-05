package containsduplicateii

func containsNearbyDuplicate(nums []int, k int) bool {
	mapping := make(map[int][]int)
	for i, num := range nums {
		n, ok := mapping[num]
		if !ok {
			mapping[num] = []int{i}
			continue
		}
		if i-n[len(n)-1] <= k {
			return true
		}
		mapping[num] = append(mapping[num], i)
	}
	return false
}
