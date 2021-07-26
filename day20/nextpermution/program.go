package nextpermution

func nextPermutation(nums []int) {
	nlen := len(nums)
	if nlen <= 1 {
		return
	}

	k := -1
	for i := nlen - 2; i > 0; i-- {
		if nums[i] <= nums[i+1] {
			k = i
			break
		}
	}
	if k == -1 {
		reverse(nums)
		return
	}
	l := 0
	for j := nlen - 1; j > k; j-- {
		if nums[j] > nums[k] {
			l = j
			nums[k], nums[l] = nums[l], nums[k]
			reverse(nums[k:])
			break
		}
	}
	// in case an arrangement is impossible
}

func reverse(nums []int) {

	if len(nums) == 0 {
		return
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
