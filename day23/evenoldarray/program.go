package evenoldarray

func Solution(arr []int) []int {

	l := len(arr)
	if l < 3 {
		return arr
	}
	start, end := 0, l-1
	for start < end {
		rest := arr[start] % 2
		if start%2 == rest {
			start++
		}
		restEnd := arr[end] % 2
		preEnd := end
		shouldSwapPreEnd := false
		for start < end {
			if restEnd == end%2 {
				end--
				continue
			}
			if restEnd == rest {
				preEnd = end
				shouldSwapPreEnd = true
				end -= 1
				restEnd = arr[end] % 2
				continue
			}
			break
		}
		if shouldSwapPreEnd {
			swap(arr, end, preEnd)
			end--
			continue
		}
		swap(arr, start, end)
		end--
		start++

	}
	return arr
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
