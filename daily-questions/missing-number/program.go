package missingnumber

func missingNumber(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	total := 0
	for i := 0; i <= len(nums); i++ {
		total += i
	}
	return total - sum
}
