package find_all_duplicates_in_an_array

//func findDuplicates(nums []int) []int {
//	freq := make(map[int]int, len(nums))
//	result := []int{}
//	for _, n := range nums {
//		if freq[n] == 0 {
//			freq[n]++
//		} else if freq[n] == 1 {
//			result = append(result, n)
//			freq[n] = 2
//		}
//	}
//	return result
//}

func findDuplicates(nums []int) []int {
	temp := 0
	result := make([]int, 0, len(nums)-1)
	for _, n := range nums {
		if n < 0 {
			continue
		}
		temp = nums[n-1]
		if nums[n-1] == -1 {
			result = append(result, n)
			nums[n-1] = -2
		} else if nums[n-1] == -2 {
			continue
		} else {
			nums[n-1] = -1
			nums[temp-1] = -1
		}
	}
	return result
}

// 1 2 ,4,4
