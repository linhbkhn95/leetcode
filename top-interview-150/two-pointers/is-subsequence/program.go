package issubsequence

func isSubsequence(s string, t string) bool {
	i, j, ls, lt := 0, 0, len(s), len(t)

	for i < ls && j < lt {
		if s[i] != t[j] {
			j++
			continue
		}
		i++
		j++
	}
	return i == ls
}
