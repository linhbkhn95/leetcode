package romantointeger

func romanToInt(s string) int {
	result := 0
	l := len(s)
	for i := 0; i < l; i++ {
		switch s[i] {
		case 'M':
			{
				result += 1000
			}
		case 'D':
			{
				result += 500

			}
		case 'C':
			{
				if i < l-1 && s[i+1] == 'M' {
					result += 900
					i += 1
				} else if i < l-1 && s[i+1] == 'D' {
					result += 400
					i += 1
				} else {
					result += 100
				}
			}
		case 'L':
			{
				result += 50
			}
		case 'X':
			{
				if i < l-1 && s[i+1] == 'C' {
					result += 90
					i += 1
				} else if i < l-1 && s[i+1] == 'L' {
					result += 40
					i += 1
				} else {
					result += 10
				}

			}

		case 'V':
			{
				result += 5
			}

		case 'I':
			{
				if i < l-1 && s[i+1] == 'X' {
					result += 9
					i += 1
				} else if i < l-1 && s[i+1] == 'V' {
					result += 4
					i += 1
				} else {
					result += 1
				}
			}
		}
	}
	return result
}
