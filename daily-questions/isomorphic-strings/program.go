package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tReplacedChar := make(map[byte]byte, 0)
	sReplacedChar := make(map[byte]byte, 0)

	for i := range s {
		tChar, ok := tReplacedChar[s[i]]
		if !ok {
			_, ok := sReplacedChar[t[i]]
			if ok {
				return false
			}

			sReplacedChar[t[i]] = s[i]
			tReplacedChar[s[i]] = t[i]
		} else {
			if tChar != t[i] {
				return false
			}
			// sChar, ok := sReplacedChar[tChar]
			// if !ok {
			// 	return false
			// }
			// if sChar != s[i] {
			// 	return false
			// }
		}
	}
	return true
}

// for
// baa
