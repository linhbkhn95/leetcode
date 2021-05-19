package minimumnumberofcoins

func Solution(coins []int) int {
	if len(coins) < 2 {
		return len(coins)
	}
	return doSolution(coins, 0, len(coins), 100, 0)
}

func doSolution(coins []int, start, len, currentResult, currentRevesed int) int {
	if currentResult < currentRevesed {
		return currentResult
	}
	preIndex := start
	for i := start + 1; i < len; i++ {
		if coins[i] != coins[preIndex] {
			preIndex = i
			continue
		}
		if preIndex == 0 {
			l1 := doSolution(coins, i, len, currentResult, currentRevesed+1)
			l2 := doSolution(coins, i+1, len, l1, currentRevesed+1)
			return min(l1, l2)
		}
		currentRevesed += 1

		if i+1 < len && coins[i] != coins[i+1] {
			currentRevesed += 1
		}
		return doSolution(coins, i+1, len, currentResult, currentRevesed)
	}
	return currentRevesed

}

func min(firstN, secondN int) int {
	if firstN > secondN {
		return secondN
	}
	return firstN
}
