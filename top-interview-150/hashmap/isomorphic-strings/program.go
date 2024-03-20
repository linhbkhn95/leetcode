package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	replaceChar := make(map[byte]string)
	str := ""
	replaceCharT := make(map[byte]struct{})

	for i, c := range t {
		replaceC, ok := replaceChar[s[i]]
		if !ok {
			_, ok := replaceCharT[t[i]]
			if !ok {
				str += string(c)
				replaceChar[s[i]] = string(c)
				replaceCharT[t[i]] = struct{}{}
			} else {
				return false
			}
		} else {
			str += replaceC
		}
	}
	return str == t
}
