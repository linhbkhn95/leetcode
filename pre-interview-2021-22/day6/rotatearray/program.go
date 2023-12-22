package rotatearray

import "fmt"

func rotate(nums []int, k int) []int {
	l := len(nums)
	k = k % l
	mapping := make(map[int]int, l)
	for i, n := range nums {
		mapping[i] = n
	}
	for i, _ := range nums {
		if i+k > l-1 {
			nums[i+k-l] = mapping[i]
		} else {
			nums[i+k] = mapping[i]
		}
	}
	// doRotate(&nums, l, k, 0, 1, nums[0])
	// fmt.Println("nums", nums)
	return nums
}

func doRotate(nums *[]int, l, k, index, start, value int) {

	if start > l {
		return
	}
	fmt.Println("doRotate", nums, k, index, start, value)

	newIndex := index + k
	if index+k > l-1 {
		newIndex = index + k - l
	}
	temp := (*nums)[newIndex]
	(*nums)[newIndex] = value
	doRotate(nums, l, k, newIndex, start+1, temp)
}


// way 1

// func rotate(nums []int, k int) []int {
// 	l := len(nums)
// 	k = k % l
// 	mapping := make(map[int]int, l)
// 	for i, n := range nums {
// 		mapping[i] = n
// 	}
// 	for i, _ := range nums {
// 		if i+k > l-1 {
// 			nums[i+k-l] = mapping[i]
// 		} else {
// 			nums[i+k] = mapping[i]
// 		}
// 	}
// 	return nums
// }