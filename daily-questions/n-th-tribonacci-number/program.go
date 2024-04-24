package main

func tribonacci(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		}
		return 1
	}
	t0, t1, t2 := 0, 1, 1
	t3 := 0
	for i := 3; i <= n; i++ {
		t3 = t1 + t2 + t0
		t0, t1, t2 = t1, t2, t3
	}
	return t3
}
