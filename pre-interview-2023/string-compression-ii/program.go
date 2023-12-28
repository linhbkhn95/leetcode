package stringcompressionii

import (
	"math"
	"sort"
)

type State struct {
	idx           int
	lastChar      byte
	lastCharCount int
	k             int
}

func getLengthOfOptimalCompression(s string, k int) int {
	m := make(map[State]int)
	return dp(0, 'a'-1, 0, k, m, s)
}

func dp(idx int, lastChar byte, lastCharCount int, k int, m map[State]int, s string) int {
	state := State{idx, lastChar, lastCharCount, k}
	if v, ok := m[state]; ok {
		return v
	}

	if k < 0 {
		return math.MaxInt
	}

	if idx == len(s) {
		return 0
	}

	var keepChar int
	deleteChar := dp(idx+1, lastChar, lastCharCount, k-1, m, s)
	if s[idx] == lastChar {
		keepChar = dp(idx+1, lastChar, lastCharCount+1, k, m, s)
		if lastCharCount == 1 || lastCharCount == 9 || lastCharCount == 99 {
			keepChar++
		}
	} else {
		keepChar = dp(idx+1, s[idx], 1, k, m, s) + 1
	}

	minv := min(keepChar, deleteChar)
	m[state] = minv
	return minv
}

func getLengthOfOptimalCompression1(s string, k int) int {
	l := len(s)
	repeatedCounts := make([]int, 0)
	lastIndex, currentCount := 0, 1
	for i := 1; i < l; i++ {
		if s[lastIndex] == s[i] {
			currentCount++
		} else {
			repeatedCounts = append(repeatedCounts, currentCount)
			lastIndex = i
			currentCount = 1
		}
	}
	repeatedCounts = append(repeatedCounts, currentCount)

	sort.Slice(repeatedCounts, func(i, j int) bool {
		return repeatedCounts[i] < repeatedCounts[j]
	})

	index := 0
	for index < len(repeatedCounts) {
		count := repeatedCounts[index]
		if k-count >= 0 {
			k -= count
			index++
		} else {
			break
		}
	}

	// Find the chactacter -> reduce length of string.
	// 100 -> 99
	// 1000 -> 999
	result := 0
	isMinus := false
	for index < len(repeatedCounts) {
		if repeatedCounts[index] < 10 {
			result += 2
		} else if repeatedCounts[index] < 100 {
			result += 3
		}
		if !isMinus && repeatedCounts[index] >= 10 && (repeatedCounts[index]-k)/10 < 1 {
			isMinus = true
			result -= 1
		}
		if !isMinus && repeatedCounts[index] >= 100 && (result-k)/100 < 1 {
			result -= 1
		}
		index++
	}

	return result
}

func getLengthOfOptimalCompression2(s string, k int) int {
	l := len(s)
	repeatedCounts := make([]int, 0)
	chars := make([]byte, 0)
	lastIndex, currentCount := 0, 1
	for i := 1; i < l; i++ {
		if s[lastIndex] == s[i] {
			currentCount++
		} else {
			chars = append(chars, s[lastIndex])
			repeatedCounts = append(repeatedCounts, currentCount)
			lastIndex = i
			currentCount = 1
		}
	}
	repeatedCounts = append(repeatedCounts, currentCount)
	chars = append(chars, s[lastIndex])

	return try(len(repeatedCounts), k, 0, repeatedCounts, chars)
}

func try(l, k int, deletedIndex int, repeatedCount []int, chars []byte) int {
	if k == 0 {
		return countLen(repeatedCount, chars)
	}
	if deletedIndex > l-1 {
		return 100
	}
	repeatedCount[deletedIndex] -= 1
	count1 := try(l, k-1, deletedIndex, repeatedCount, chars)
	repeatedCount[deletedIndex] += 1
	count2 := try(l, k, deletedIndex+1, repeatedCount, chars)

	return min(count1, count2)
}

func countLen(repeatedCount []int, chars []byte) int {
	result := 0
	lastIndex := 0
	charCount := repeatedCount[0]
	for i := 1; i < len(repeatedCount); i++ {
		if repeatedCount[i] == 0 {
			continue
		}
		if chars[i] == chars[lastIndex] {
			charCount += repeatedCount[i]
			continue
		}
		result += calcRealLen(charCount)
		lastIndex = i
		charCount = repeatedCount[i]

	}
	result += calcRealLen(charCount)
	return result
}

func calcRealLen(charCount int) int {
	count := 0
	switch {
	case charCount == 1:
		{
			count = 1
		}
	case charCount < 10:
		count = 2

	case charCount < 100:
		count = 3
	case charCount == 100:
		count = 3
	}
	return count

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
