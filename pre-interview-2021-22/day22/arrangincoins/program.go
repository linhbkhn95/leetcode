package arrangincoins

func arrangeCoins(n int) int {
	s := 0
	i := 1
	for {
		if s+i > n {
			return i - 1
		}
		s += i
		i++
	}
}
