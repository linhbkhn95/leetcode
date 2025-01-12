package checkifaparenthesesstringcanbevalid

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	balance := 0
	for i := 0; i < len(s); i++ {
		if locked[i] == '0' || s[i] == '(' {
			balance++
		} else if s[i] == ')' {
			balance--
		}
		if balance < 0 {
			return false
		}
	}

	balance = 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '0' || s[i] == ')' {
			balance++
		} else if s[i] == '(' {
			balance--
		}
		if balance < 0 {
			return false
		}
	}

	return true
}
