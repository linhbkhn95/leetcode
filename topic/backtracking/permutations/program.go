package permutations

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	choosed := make(map[int]struct{})
	backtracking(nums, []int{}, choosed, &result)
	return result
}

func backtracking(nums []int, candidates []int, choosed map[int]struct{}, result *[][]int) {

	if len(candidates) == len(nums) {
		p := make([]int, len(candidates))
		copy(p, candidates)
		(*result) = append((*result), p)
		return
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := choosed[nums[i]]; !ok {
			candidates = append(candidates, nums[i])
			choosed[nums[i]] = struct{}{}
			backtracking(nums, candidates, choosed, result)
			delete(choosed, nums[i])
			candidates = candidates[:len(candidates)-1]
		}
	}

}
