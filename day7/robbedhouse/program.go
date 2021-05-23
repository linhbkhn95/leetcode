package robbedhouse

func rob(nums []int) int {

	// use to mark point that way passed attached current sum
	mark := make(map[int]int)
	return doRob(nums, &mark, len(nums), 0, 0, 0)

}

func doRob(arr []int, mark *map[int]int, len, index, currentSum, largestSum int) int {
	if index > len-1 {
		if currentSum > largestSum {
			return currentSum
		}
		return largestSum
	}

	value, ok := (*mark)[index]
	if ok && currentSum <= value {
		return largestSum
	}

	(*mark)[index] = currentSum
	case1 := doRob(arr, mark, len, index+2, currentSum+arr[index], largestSum)
	case2 := doRob(arr, mark, len, index+1, currentSum, largestSum)
	if case1 > case2 {
		return case1
	}
	return case2
}
