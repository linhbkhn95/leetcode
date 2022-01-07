package regularexpression

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	isFirstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return (isFirstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
	}
	return isFirstMatch && isMatch(s[1:], p[1:])

}
