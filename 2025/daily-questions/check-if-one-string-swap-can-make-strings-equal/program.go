package checkifonestringswapcanmakestringsequal

func areAlmostEqual(s1 string, s2 string) bool {
	pos1, pos2 := -1, -1
	if s1 == s2 {
		return true
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if pos1 == -1 {
			pos1 = i
		} else if pos2 == -1 {
			pos2 = i
		} else {
			return false
		}

	}
	if pos2 == -1 {
		return false
	}

	return s1[pos1] == s2[pos2] && s1[pos2] == s2[pos1]
}
