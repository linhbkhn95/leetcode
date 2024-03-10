package minimum_common_value

//func getCommon(nums1 []int, nums2 []int) int {
//	mapping := make(map[int]int)
//	for _, n := range nums1 {
//		mapping[n]++
//	}
//	for _, n := range nums2 {
//		if mapping[n] > 0 {
//			return n
//		}
//	}
//	return -1
//}

func getCommon(nums1 []int, nums2 []int) int {

	index1, index2, l1, l2 := 0, 0, len(nums1), len(nums2)
	for index1 < l1 && index2 < l2 {
		if nums2[index2] == nums1[index1] {
			return nums2[index2]
		} else if nums2[index2] > nums1[index1] {
			index1++
		} else {
			index2++
		}
	}
	return -1
}
