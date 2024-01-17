package majorityelement

func majorityElement(nums []int) int {
	occurrences := make(map[int]int)
	for _, n := range nums {
		occurrences[n]++
	}
	len := len(nums)
	for n, c := range occurrences {
		if c >= len/2 {
			return n
		}
	}
	return 0
}
