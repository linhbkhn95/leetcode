package longestpalindrome

func longestPalindrome(s string) string {
	max := 0
	result := ""
	l := len(s)
	for i := 0; i < l; i++ {
		for j := l; j > i; j-- {
			if j-i < max {
				continue
			}
			ss := s[i:j]
			valid := isPalindrome(ss)
			if valid && max < j-i {
				result = ss
				max = j - i
			}
		}
	}
	return result
}

func isPalindrome(s string) bool {

	len := len(s)
	if len < 2 {
		return true
	}
	i, j := 0, len-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
