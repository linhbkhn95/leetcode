package buildarrwithstackops

func buildArray(target []int, n int) []string {
	result := make([]string, 0, n*2)
	j := 0
	for i := 1; i <= n; {
		if j == len(target) {
			return result
		}
		if i != target[j] {
			result = append(result, "Push", "Pop")
			i++
		} else {
			result = append(result, "Push")
			j++
			i++
		}

	}
	return result
}
