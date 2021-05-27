package sortcolor

func sortColors(nums []int) {

	quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := (left + right) / 2

	index := patition(nums, left, right, pivot)

	quicksort(nums, left, index-1)
	quicksort(nums, index, right)
}

func patition(arr []int, left, right, pivot int) int {
	for left <= right {
		for arr[left] < arr[pivot] {
			left++
		}
		for arr[right] > arr[pivot] {
			right--
		}
		if left <= right {
			swap(arr, left, right)
			left++
			right--
		}
	}
	return left
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
