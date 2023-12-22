package pow

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 1.00000 {
		return 1
	}
	if x == -1.00000 {
		if n < 0 {
			return 1
		}
		if n%2 == 1 {
			return -1
		}
		return 1
	}
	result := 1.00000

	if n > 0 {
		for n > 0 {
			result *= x
			n--
		}
	} else {
		for n < 0 {
			result *= x
			n++
		}
		result = 1 / result
	}
	return result
}
