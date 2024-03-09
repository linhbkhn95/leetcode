package count_elements_with_maximum_frequency

func maxFrequencyElements(nums []int) int {
	var result int
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	maxFreq := 0
	for _, f := range freq {
		maxFreq = max(maxFreq, f)
	}
	for _, f := range freq {
		if maxFreq == f {
			result += f
		}
	}
	return result
}
