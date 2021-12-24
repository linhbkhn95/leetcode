package program

func Solution(A []int) int {

	maxLen := 0
	currentLen := 1
	l := len(A)
	if l < 3 {
		return l
	}
	for i := 0; i < l-1; i++ {
		if maxLen >= (l - i) {
			return maxLen
		}
		visited := make(map[int]interface{}, 2)
		j := i
		currentLen = 0
		visitedCount := 1
		visited[A[j]] = nil

		for j < l {
			_, ok := visited[A[j]]
			if ok && visitedCount == 2 || visitedCount == 1 {
				if !ok {
					visitedCount++
					visited[A[j]] = nil
				}

				currentLen++
				j++
			} else if !ok {
				maxLen = max(maxLen, currentLen)
				break
			}
		}
		maxLen = max(maxLen, currentLen)
	}

	return max(maxLen, currentLen)
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}
