package cleardigits

var digits = map[byte]bool{
	'9': true,
	'8': true,
	'7': true,
	'6': true,
	'5': true,
	'4': true,
	'3': true,
	'2': true,
	'1': true,
	'0': true,
}

func clearDigits(s string) string {
	result := []byte{}
	for i := 0; i < len(s); i++ {
		if digits[s[i]] {
			result = result[:len(result)-1]
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
