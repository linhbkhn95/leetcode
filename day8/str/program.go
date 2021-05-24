package str

func strStr(haystack string, needle string) int {
	l := len(needle)
	lStack := len(haystack)
	if l == 0 {
		return 0
	}
	if l > lStack {
		return -1
	}
	for i, _ := range haystack {
		if haystack[i] == needle[0] && l+i <= lStack && haystack[l-1+i] == needle[l-1] {
			if haystack[i:l+i] == needle {
				return i
			}
		}
	}
	return -1
}
