package main

func maxLength(arr []string) int {
	maxLen := 0
	dp(arr, len(arr), 0, false, [26]int{}, 0, &maxLen)
	dp(arr, len(arr), 0, true, [26]int{}, 0, &maxLen)

	return maxLen
}

func dp(arr []string, l, start int, shouldConcat bool, mask [26]int, currentLen int, maxLen *int) {
	if start == l {
		*maxLen = max(*maxLen, currentLen)
		return
	}
	maskCopy := [26]int{}
	copy(maskCopy[:], mask[:])
	if shouldConcat {
		if canConcate(&mask, arr[start]) {
			dp(arr, l, start+1, false, mask, currentLen+len(arr[start]), maxLen)
			dp(arr, l, start+1, true, mask, currentLen+len(arr[start]), maxLen)
		}
	} else {
		dp(arr, l, start+1, false, maskCopy, currentLen, maxLen)
		dp(arr, l, start+1, true, maskCopy, currentLen, maxLen)

	}

}

func canConcate(mask *[26]int, c string) bool {
	for i := range c {
		if mask[c[i]-'a'] == 1 {
			return false
		}
		mask[c[i]-'a']++
	}
	return true
}
