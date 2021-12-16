package bit01operation

func Solution(s string) int32 {
	stringBytes := []byte(s)
	result := 0
	i := len(s) - 1
	k := 0
	for k < len(s)-1 {
		if stringBytes[k] == '0' {
			k++
		} else {
			break
		}
	}

	for i >= k {
		if stringBytes[i] == '0' {
			result++
			i--
		} else {
			result++
			stringBytes[i] = '0'
			if i == k {
				break
			}
		}
	}
	return int32(result)
}
