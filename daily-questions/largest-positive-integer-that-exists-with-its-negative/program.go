package largestpositiveintegerthatexistswithitsnegative

func findMaxK(nums []int) int {
	result := 0
	mapper := make(map[int]interface{})
	for _, n := range nums {
		mapper[n] = struct{}{}
	}
	for _, n := range nums {
		if n > 0 {
			if _, ok := mapper[-n]; ok {
				result = max(result, n)
			}
		}
	}
	if result == 0 {
		return -1
	}
	return result
}
