package find_the_pivot_integer

func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}
	start, end := 2, n-1
	sLeft, sRight := 1, n
	for start <= end {
		if start == end && sLeft == sRight {
			return start
		}
		if sLeft < sRight {
			sLeft += start
			start++
		} else {
			sRight += end
			end--
		}
	}
	return -1
}
