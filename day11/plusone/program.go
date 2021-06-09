package plusone

func plusOne(digits []int) []int {

	remain := 1
	for i := len(digits) - 1; i > -1; i-- {
		digits[i] += remain
		if digits[i] == 10 {
			digits[i] = 0
			remain = 1
		} else {
			remain = 0
		}
	}
	if remain == 1 && digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
