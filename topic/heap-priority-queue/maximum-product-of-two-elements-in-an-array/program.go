package maximumproductoftwoelementsinanarray

func maxProduct(nums []int) int {
	max1, max2 := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		if nums[i] > max1 {
			if max1 > max2 {
				max2 = nums[i]
			} else {
				max1 = nums[i]
			}
		} else if nums[i] > max2 {
			if max2 > max1 {
				max1 = nums[i]
			} else {
				max2 = nums[i]
			}
		}
	}
	return (max1 - 1) * (max2 - 1)
}
