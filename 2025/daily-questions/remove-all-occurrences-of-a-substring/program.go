package removealloccurrencesofasubstring

import "container/list"

func removeOccurrences(s string, part string) string {
	startPartIndexes := list.New()
	lPart := len(part)
	i := 0
	var substring string
	for i+lPart <= len(s) {
		if s[i] == part[0] {
			if i+lPart > len(s) {
				break
			}
			substring = s[i : i+lPart]
			if substring == part {
				s = s[:i] + s[i+lPart:]
				for startPartIndexes.Len() > 0 {
					i = startPartIndexes.Back().Value.(int)
					if i+lPart <= len(s) && s[i:i+lPart] == part {
						s = s[:i] + s[i+lPart:]
						startPartIndexes.Remove(startPartIndexes.Back())
					} else {
						break
					}
				}
			} else {
				startPartIndexes.PushBack(i)
				i++
			}
		} else {
			i++
		}
	}

	return s
}
