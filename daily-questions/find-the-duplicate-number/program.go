package find_the_duplicate_number

func findDuplicates(nums []int) []int {
	output := make([]int, 0, len(nums)-1)
	for _, num := range nums {
		idx := abs(num)
		if nums[idx-1] < 0 {
			output = append(output, idx)
		} else {
			nums[idx-1] *= -1
		}
	}
	return output
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// 1,2,3,4
// 3,3,2,1
// 1,2,4,3,3 1->2>4->3->3
// 1,3,4,2,2  1->3->2->4->2
