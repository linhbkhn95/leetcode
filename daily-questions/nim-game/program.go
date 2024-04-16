package nimgame

func canWinNim(n int) bool {
	if n < 3 {
		return true
	}
	return n%4 != 0
}
