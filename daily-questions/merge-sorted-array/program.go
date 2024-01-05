package mergesortedarray

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// i, j := 0, 0
	// temp := 0
	// for i < m && j < n {
	// 	if nums1[i] < nums2[j] {
	// 		for nums1[i] < nums2[j] && nums1[i+1] != 0 && nums1[i+1] > nums2[j] {
	// 			nums1[i] = nums2[j]
	// 			temp = nums2[i]
	// 			nums2[j] = temp
	// 			i++
	// 		}
	// 	} else {
	// 		nums1[i] = nums2[j]
	// 		temp = nums2[i]
	// 		nums2[j] = temp
	// 		j++
	// 		i++
	// 	}
	// }
	nums1 = append(nums1[:n], nums2...)
	sort.Ints(nums1)
}

// 2456000 125 145

//137    457   23
