package permutationsii

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	choosed := make(map[int]struct{})
	backtracking(nums, []int{}, choosed, &result)
	return result
}

func backtracking(nums []int, candidates []int, choosed map[int]struct{}, result *[][]int) {
	if len(candidates) == len(nums) {
		if isExisted(candidates, result) {
			return
		}
		p := make([]int, 0, len(candidates))
		copy(p, candidates)
		(*result) = append((*result), candidates)
	}
	for i := range nums {
		if _, ok := choosed[i]; !ok {
			candidates = append(candidates, nums[i])
			choosed[i] = struct{}{}
			backtracking(nums, candidates, choosed, result)
			delete(choosed, i)
			candidates = candidates[:len(candidates)-1]
		}
	}
}

func isExisted(candicate []int, result *[][]int) bool {
	for _, arr := range *result {
		if isSameArr(arr, candicate) {
			return true
		}
	}
	return false
}

func isSameArr(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
