package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	// offsetChar1 := make(map[byte][]int, 0)
	// for i := range text1 {
	// 	offsetChar1[text1[i]] = append(offsetChar1[text1[i]], i)
	// }
	// var result int
	// backtracking(text2, len(text2), 0, -1, 0, &result, offsetChar1)
	// return result
	return dpSolutionWithoutRecursive(text1, text2)
	return dp(text1, text2, 0, 0, map[key]int{})
}

func dpSolutionWithoutRecursive(text1, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]

}

type key struct {
	i, j int
}

func dp(text1, text2 string, i, j int, footprint map[key]int) int {
	if i == len(text1) || j == len(text2) {
		return 0
	}
	k := key{
		i: i,
		j: j,
	}
	if v, ok := footprint[k]; ok {
		return v
	}
	var maxLen int
	if text1[i] == text2[j] {
		maxLen = 1 + dp(text1, text2, i+1, j+1, footprint)
	} else {
		maxLen = max(dp(text1, text2, i, j+1, footprint), dp(text1, text2, i+1, j, footprint))
	}
	footprint[k] = maxLen
	return maxLen
}
