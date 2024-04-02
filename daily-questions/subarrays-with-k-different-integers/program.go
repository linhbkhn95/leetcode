package subarrays_with_k_different_integers

func subarraysWithKDistinct(nums []int, k int) int {
	freqs := make(map[int]int)
	result := 0
	start, end, l := 0, 0, len(nums)
	firstOffsetForEnd := 0
	for start < l && end < l {
		freqs[nums[end]]++
		if len(freqs) == k {
			end++
		} else if len(freqs) > k {
			delete(freqs, nums[end])
			freqs[nums[start]]--
			if freqs[nums[start]] == 0 {
				delete(freqs, nums[start])
			} else {
				result += end - firstOffsetForEnd - 1
				start += 1
			}
		} else {
			firstOffsetForEnd = end
			end++
		}
	}
	return result
}
