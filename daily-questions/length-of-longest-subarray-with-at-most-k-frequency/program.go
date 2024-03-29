package lengthoflongestsubarraywithatmostkfrequency

func maxSubarrayLength(nums []int, k int) int {
	freqs := make(map[int]int)
	start, end, l := 0, 0, len(nums)
	maxLen := 0
	for end < l {
		num := nums[end]
		f, ok := freqs[num]
		if !ok || f < k {
			freqs[num]++
			end++
			continue
		}
		maxLen = max(maxLen, end-start)
		i := start
		for nums[i] != num {
			freqs[nums[i]]--
			i++
		}
		start = i + 1
		end++
	}
	return max(maxLen, end-start)
}
