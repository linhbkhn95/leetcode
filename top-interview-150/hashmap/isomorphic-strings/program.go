package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	replaceChar := make(map[byte]struct{})
	replaceCharT := make(map[byte]struct{})

	for i := range t {
		_, ok := replaceChar[s[i]]
		if !ok {
			_, ok := replaceCharT[t[i]]
			if !ok {
				replaceChar[s[i]] = struct{}{}
				replaceCharT[t[i]] = struct{}{}
			} else {
				return false
			}
		}
	}
	return true
}
