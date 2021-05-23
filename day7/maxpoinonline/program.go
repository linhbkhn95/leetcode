package maxpoinonline

type line struct {
	finish bool
	maxLen int
}

type key struct {
	a float32
	y int
}

func maxPoints(points [][]int) int {
	mark := make(map[key]*line)
	maxLen := 1
	len := len(points)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			maxLen = max(maxLen, solution(points, mark, key{}, len, i, j, 1))
		}
	}
	return maxLen
}

// solution
func solution(points [][]int, mark map[key]*line, k key, len, start, index, currentLen int) int {
	if index > len-1 {

		return currentLen

	}
	kA := key{}
	if (points[index][1] - points[start][1]) != 0 {
		kA.a = float32(points[index][0]-points[start][0]) / float32(points[index][1]-points[start][1])
	} else {
		kA.y = points[index][1]
	}
	if currentLen == 1 {
		k = kA
	}
	value, ok := mark[k]
	// is'nt belong to line
	if !ok && currentLen > 1 {
		return solution(points, mark, k, len, start, index+1, currentLen)
	}

	if ok && value.finish && value.maxLen > (len-index+currentLen) && equal(k, kA) {
		return value.maxLen
	}
	if ok && currentLen > 1 && !equal(k, kA) {
		return solution(points, mark, k, len, start, index+1, currentLen)
	}
	currentLen += 1

	if !ok {
		mark[k] = &line{
			finish: false,
			maxLen: currentLen,
		}
	} else {
		if !value.finish {
			mark[k].maxLen = currentLen
		}

	}
	if index == len-1 {
		if ok && value.finish {
			mark[k].maxLen = max(currentLen, value.maxLen)
		}
		mark[k].finish = true
		return max(currentLen, mark[k].maxLen)
	}

	return solution(points, mark, k, len, start, index+1, currentLen)

}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func equal(k1, k2 key) bool {
	if k1.a != k2.a {
		return false
	}
	if k1.y != k2.y {
		return false
	}
	return true
}
