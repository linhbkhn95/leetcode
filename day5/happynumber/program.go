package happynumber

func isHappy(n int) bool {
	if n < 1 {
		return false
	}
	return do(n)
}

func do(n int) bool {
	mark := make(map[int]bool)
	for {

		mark[n] = true
		n = calcSum(n)
		if n == 1 {
			return true
		}
		if _, ok := mark[n]; ok {
			return false
		}
	}
}

func calcSum(n int) int {
	s := 0
	for n >= 10 {
		remain := n % 10
		s += remain * remain
		n = n / 10
	}
	s += n * n
	return s
}
