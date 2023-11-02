package sortbybits

func sortByBits(arr []int) []int {
	countMapping := make(map[int]int, len(arr))
	quickSort(arr, countMapping, 0, len(arr)-1)
	start, end := 0, 0
	for i := 0; i < len(arr)-1; i++ {
		if countMapping[arr[i]] == countMapping[arr[i+1]] {
			end++
		} else {
			customQuickSort(arr, start, end)
			start = i + 1
			end = i + 1
		}
	}
	return arr
}

func customQuickSort(arr []int, left, right int) {
	if left < right {

		pivot := (left + right) / 2

		leftCondition := func(i, pivot int) bool {
			return arr[i] < arr[pivot]
		}
		rightCondition := func(i, pivot int) bool {
			return arr[i] > arr[pivot]
		}
		index := partition(arr, leftCondition, rightCondition, left, right, pivot)
		customQuickSort(arr, left, index-1)
		customQuickSort(arr, index+1, right)
	}
}

func quickSort(arr []int, mapping map[int]int, left, right int) {
	if left < right {

		pivot := (left + right) / 2

		leftCondition := func(i, pivot int) bool {
			return count(mapping, arr[i]) < count(mapping, arr[pivot])
		}
		rightCondition := func(i, pivot int) bool {
			return count(mapping, arr[i]) > count(mapping, arr[pivot])
		}
		index := partition(arr, leftCondition, rightCondition, left, right, pivot)
		quickSort(arr, mapping, left, index-1)
		quickSort(arr, mapping, index+1, right)
	}
}

func partition(arr []int, leftCondition func(left, pivot int) bool, rightCondition func(right, pivot int) bool, left, right, pivot int) int {
	for left <= right {
		for leftCondition(left, pivot) {
			left++
		}
		for rightCondition(right, pivot) {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

func count(mapping map[int]int, n int) int {
	count := 0
	number := n
	for number > 1 {
		c, ok := mapping[number]
		if ok {
			mapping[n] = count + c
			return count + c
		}
		if number%2 == 1 {
			count++
		}
		number = number / 2

	}
	if number == 1 {
		count++
	}
	mapping[n] = count
	return mapping[n]
}
