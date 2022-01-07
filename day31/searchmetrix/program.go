package program

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	for _, arr := range matrix {
		if arr[0] <= target && arr[n-1] >= target {
			return binarySearch(arr, n, target)
		}
		continue
	}
	return false
}

func binarySearch(arr []int, l, target int) bool {

	start := 0
	end := l - 1

	for start <= end {
		middle := (start + end) / 2

		if arr[middle] == target {
			return true
		} else if arr[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return false
}
