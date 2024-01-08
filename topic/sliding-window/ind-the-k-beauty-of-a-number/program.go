package indthekbeautyofanumber

import "strconv"

func divisorSubstrings(num int, k int) int {
	sNumber := strconv.Itoa(num)
	result := 0
	for i := 0; i < len(sNumber)-k+1; i++ {
		substring := sNumber[i : i+k]
		n, _ := strconv.Atoi(substring)
		if n == 0 {
			continue
		}
		if num%n == 0 {
			result++
		}
	}
	return result
}
