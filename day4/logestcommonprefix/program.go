package logestcommonprefix

func longestCommonPrefix(strs []string) string {

	min := min(strs)
	stop := false
	i := 0
	for i < len(min) {
		for _, str := range strs {
			if str[i] != min[i] {
				stop = true
			}
		}
		if stop {
			break
		}
		i++
	}
	return min[:i]
}

func min(strs []string) string {

	result := strs[0]
	for _, str := range strs {
		if len(str) < len(result) {
			result = str
		}
	}
	return result
}
