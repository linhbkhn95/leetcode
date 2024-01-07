package partitionarrayformaximumsum

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr))

	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		maxVV, maxV := 0, 0
		for j := i; j >= max(i-k+1, 0); j-- {
			maxV = max(arr[j], maxV)
			calc := (i - j + 1) * maxV
			if j > 0 {
				maxVV = max(maxVV, calc+dp[j-1])
			} else {
				maxVV = max(maxVV, calc)
			}
		}
		dp[i] = maxVV
	}
	return dp[len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	// Implement your solution here

	mapping := make(map[int]struct{})
	for _, num := range A {
		if num > 0 {
			mapping[num] = struct{}{}
		}
	}
	for i := 1; i < 1000000; i++ {
		if _, ok := mapping[i]; !ok {
			return i
		}
	}
	return -1
}
