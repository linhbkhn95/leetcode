package ms

func removeDuplicateLetters(s string) string {
	maxIndexChar := make(map[byte]int)
	for i := range s {
		maxIndexChar[s[i]] = i
	}
	result := make([]byte, 0, len(s))
	existedChar := make(map[byte]bool)
	i := 0
	for i < len(s) {
		c := s[i]
		if len(result) == 0 {
			existedChar[c] = true
			result = append(result, c)
			i++
			continue
		}
		if c == result[len(result)-1] || existedChar[c] {
			i++
			continue
		}
		if c > result[len(result)-1] {
			result = append(result, c)
			existedChar[c] = true
			i++
			continue
		}
		maxOffset := maxIndexChar[result[len(result)-1]]
		if maxOffset <= i {
			result = append(result, c)
			existedChar[c] = true
			continue
		}
		delete(existedChar, result[len(result)-1])
		result = result[:len(result)-1]
	}
	return string(result)
}

// cbabc ->
// need result []byte
// need appeared map[byte]int
// loop -> 0 -> n-1
// ex c -> result
//    b -> result
//    a -> result  current cba
//    b -> cab -> result = min
//    c -> abc -> abc and cab
