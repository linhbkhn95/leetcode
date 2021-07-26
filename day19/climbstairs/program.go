package climbstairs

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	arr := make([]int, n)
	arr[0], arr[1] = 1, 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n-1]
}
