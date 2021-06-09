package sortcolor

func sortColors(nums []int) {

	if len(nums) <= 1 {
		return
	}
	var start, end = 0, len(nums) - 1
	for i := 0; i <= end; i++ {
		// swap current to end
		if nums[i] == 2 && i != end {
			nums[end], nums[i] = nums[i], nums[end]
			end--
			i--
		} else if nums[i] == 0 && i != start {
			nums[start], nums[i] = nums[i], nums[start]
			start++
		}
	}

}

// func quicksort(nums []int, left, right int) {
// 	if left >= right {
// 		return
// 	}

// 	pivot := nums[(left+right)/2]
// 	index := patition(nums, left, right, pivot)
// 	quicksort(nums, left, index-1)
// 	quicksort(nums, index, right)
// }

// func patition(arr []int, left, right, pivot int) int {
// 	for left <= right {
// 		for arr[left] < pivot {
// 			left++
// 		}
// 		for arr[right] > pivot {
// 			right--
// 		}
// 		if left <= right {
// 			swap(arr, left, right)
// 			left++
// 			right--
// 		}
// 	}
// 	return left
// }

// func swap(arr []int, i, j int) {
// 	temp := arr[i]
// 	arr[i] = arr[j]
// 	arr[j] = temp
// }
