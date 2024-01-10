package defusethebomb

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	for i := range code {
		sum := 0
		if k > 0 {
			count := 0
			for j := i + 1; j < len(code); j++ {
				if count == k {
					break
				}
				count++
				sum += code[j]
			}
			for u := 0; u < k-count; u++ {
				sum += code[u]
			}
		} else if k < 0 {
			count := 0

			for j := i - 1; j >= 0; j-- {
				if count == k*-1 {
					break
				}
				count++
				sum += code[j]
			}
			for i := 0; i < k*-1-count; i++ {
				sum += code[len(code)-1-i]
			}
		}
		result[i] = sum
	}
	return result
}
