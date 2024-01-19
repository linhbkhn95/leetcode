package climbingstairs

func climbStairs(n int) int {
	return dp(n)
}

func dp(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1, n2 := 1, 2
	var result int
	for i := 3; i <= n; i++ {
		result = n1 + n2
		n1 = n2
		n2 = result
	}
	return result
}
