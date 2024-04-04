package reversewordsinastring

func reverseWords(s string) string {
	// start, end := 0, 0
	return ""
}

func trimSpace(s string) string {
	i := 0
	shouldIgnoreSpace := false
	for i < len(s) {
		if s[i] == ' ' {
			if i == 0 {
				s = s[i+1:]
			}
		}
		if i == 0 && s[i] == ' ' {
			s = s[i+1:]
		}
		if i > 0 && shouldIgnoreSpace && s[i] == ' ' {
			s = s[i-1 : i+1]
		}
		i++
	}
	return s
}

/*
Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


//

*/
