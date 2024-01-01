package maximumproductsubarray

func maxProduct(nums []int) int {

	maxV := nums[0]
	currentProduct := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			currentProduct = 1
			maxV = max(maxV, 0)
			continue
		}
		currentProduct *= nums[i]
		maxV = max(currentProduct, maxV)
	}
	currentProduct = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			currentProduct = 1
			continue
		}
		currentProduct *= nums[i]
		maxV = max(currentProduct, maxV)
	}
	return maxV
}

func maxProduct1(nums []int) int {
	// dps := make([]int, len(nums))
	// dps[0] = nums[0]
	// maxV := nums[0]
	// for i := 1; i < len(nums); i++ {
	// 	dps[i] = max(nums[i], dps[i-1]*nums[i])
	// 	maxV = max(dps[i], maxV)
	// }
	// return maxV
	maxV := nums[0]
	for i := 0; i < len(nums); i++ {
		currentProduct := nums[i]
		maxV = max(maxV, currentProduct)
		for j := i + 1; j < len(nums); j++ {
			currentProduct *= nums[j]
			maxV = max(maxV, currentProduct)
		}
	}
	return maxV
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
