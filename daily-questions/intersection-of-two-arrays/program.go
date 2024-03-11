package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	uniqueN1 := make(map[int]interface{})
	for _, n := range nums1 {
		uniqueN1[n] = struct{}{}
	}
	uniqueInsertion := make(map[int]interface{})
	for _, n := range nums2 {
		if _, ok := uniqueN1[n]; ok {
			uniqueInsertion[n] = struct {
			}{}
		}
	}
	result := make([]int, 0, len(uniqueInsertion))
	for n, _ := range uniqueInsertion {
		result = append(result, n)
	}
	return result
}
