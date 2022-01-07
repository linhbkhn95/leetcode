package addbinary

import "fmt"

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	additional := "0"
	result := ""

	i, j := la-1, lb-1
	for i > -1 && j > -1 {
		if a[i]+b[j] == '0'+'0' {
			result = additional + result
			additional = "0"
		} else if a[i]+b[j] == '0'+'1' {
			if additional == "0" {
				result = "1" + result
				additional = "0"
			} else {
				result = "0" + result
				additional = "1"
			}
		} else {
			result = additional + result
			additional = "1"
		}
		i--
		j--
	}
	check, k := a, i

	if lb > la {
		check, k = b, j
	}
	for k > -1 {
		if additional == "1" && check[k] == '1' {
			result = "0" + result
			additional = "1"
		} else {
			if additional == "1" {
				result = "1" + result
			}
			break
		}
		k--
	}
	if k > -1 {
		result = check[:k+1] + result
	} else if additional == "1" {
		result = "1" + result
	}
	fmt.Println("result", result)
	return result
}

func Solution(s string) string {
	l := len(s)
	maxLen := 0
	result := ""
	i := 0
	for i < l {
		uniqueString, lenString := extractUniqueString(s, l, i)
		if maxLen < lenString {
			maxLen = lenString
			result = uniqueString
		}
		i++
	}
	return result
}

func extractUniqueString(s string, l, index int) (string, int) {
	flag := make(map[byte]interface{})
	end := index
	for j := index; j < l; j++ {
		if _, ok := flag[s[j]]; !ok {
			end++
			flag[s[j]] = nil
		} else {
			break
		}
	}
	return s[index:end], end - index
}
