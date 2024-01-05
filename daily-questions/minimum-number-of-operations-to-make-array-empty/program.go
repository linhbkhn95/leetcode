package minimumnumberofoperationstomakearrayempty

func minOperations(nums []int) int {
	appearedCount := make(map[int]int)
	for _, n := range nums {
		appearedCount[n]++
	}

	result := 0

	for _, count := range appearedCount {
		if count == 1 {
			return -1
		}
		remain := count / 2
		redunt := count % 2

		rreamin := remain / 3
		rredunt := remain % 3
		result += rreamin*2 + max(redunt, rredunt)

	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
