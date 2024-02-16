package main

func rearrangeArray(nums []int) []int {

	positiveArr := make([]int, 0, len(nums)/2)
	negativeArr := make([]int, 0, len(nums)/2)
	for _, num := range nums {
		if num > 0 {
			positiveArr = append(positiveArr, num)
		} else {
			negativeArr = append(negativeArr, num)
		}
	}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = positiveArr[i/2]
		} else {
			nums[i] = negativeArr[(i-1)/2]
		}
	}
	return nums
}
