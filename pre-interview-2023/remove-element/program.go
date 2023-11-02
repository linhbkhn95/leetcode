package removeelement

func removeElement(nums []int, val int) int {
	l := len(nums)
	indexes := []int{}
	for i, num := range nums {
		if val == num {
			l--
			indexes = append(indexes, i)
		}
	}

	for i, index := range indexes {
		nums = append(nums[:index-i], nums[index-i+1:]...)
	}
	return l
}
