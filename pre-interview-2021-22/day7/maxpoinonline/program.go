package maxpoinonline

import "fmt"

type (
	line struct {
		finish     bool
		maxLen     int
		points     []*[]int
		sets       map[*[]int]int
		startIndex int
	}

	key struct {
		a float32
		y int
	}
	counter struct {
		value int
	}
)

func maxPoints(points [][]int) int {
	mark := make(map[key]*line)
	maxLen := 1
	len := len(points)
	count := &counter{
		value: 0,
	}
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			maxLen = max(maxLen, solution(count, points, mark, key{}, len, i, j, 1))
		}
	}
	fmt.Println("counter", count)
	return maxLen
}

// solution
func solution(count *counter, points [][]int, mark map[key]*line, k key, pointsLen, start, index, currentLen int) int {
	count.value += 1
	if index > pointsLen-1 {
		return currentLen
	}
	kA := getKey(points[start], points[index])
	if currentLen == 1 {
		k = kA
	}
	value, ok := mark[k]

	// is'nt belong to line
	if !ok && currentLen > 1 {
		return solution(count, points, mark, k, pointsLen, start, index+1, currentLen)
	}
	if ok && value.finish && value.maxLen > (pointsLen-index+currentLen) && equal(k, kA) {
		return value.maxLen
	}
	//cuting branch that result passed.
	if ok && value.finish && equal(k, kA) && equal(k, getKey(points[value.startIndex], points[index])) {
		offset := value.sets[&points[start]]
		for i := offset; i > -1; i-- {
			if _, ok := value.sets[&points[i]]; ok {
				fmt.Println("cut,1", k, pointsLen, start, index, currentLen)
				return value.maxLen
			}
		}
		if offset, ok := value.sets[&points[index]]; ok {
			for i := offset; i > -1; i-- {
				if _, ok := value.sets[&points[i]]; ok {
					fmt.Println("cut2,", k, pointsLen, start, index, currentLen)
					return value.maxLen
				}
			}
		}
	}

	if ok && currentLen > 1 && !equal(k, kA) {
		return solution(count, points, mark, k, pointsLen, start, index+1, currentLen)
	}
	currentLen += 1

	if !ok {
		mark[k] = &line{
			finish: false,
			maxLen: currentLen,
			points: []*[]int{
				&points[start], &points[index],
			},
			sets: map[*[]int]int{
				&points[start]: 0,
				&points[index]: 1,
			},
			startIndex: start,
		}
	} else {
		mark[k].points = append(mark[k].points, &points[index])
		mark[k].sets[&points[index]] = len(mark[k].points) - 1
		if !value.finish {
			mark[k].maxLen = currentLen
		}

	}
	if index == pointsLen-1 {
		if ok && value.finish {
			mark[k].maxLen = max(currentLen, value.maxLen)
		}
		mark[k].finish = true
		return max(currentLen, mark[k].maxLen)
	}

	return solution(count, points, mark, k, pointsLen, start, index+1, currentLen)

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

func getKey(point1, point2 []int) key {
	kA := key{}
	if (point1[1] - point2[1]) != 0 {
		kA.a = float32(point1[0]-point2[0]) / float32(point1[1]-point2[1])
	} else {
		kA.y = point2[1]
	}

	return kA
}
