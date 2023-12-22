package powertwo

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}

	for n != 1 {
		if n%2 == 1 {
			return false
		}
		n = n / 2
	}

	return true
}
