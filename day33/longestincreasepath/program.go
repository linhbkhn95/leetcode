package program

import (
	"fmt"
	"math"
	"sort"
)

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

type Item struct {
	i   int
	j   int
	val int
}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	var arr []Item
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr = append(arr, Item{i: i, j: j, val: matrix[i][j]})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val
	})
	fmt.Println("arr", arr)
	return findLongest(arr, m*n)
}

func findLongest(arr []Item, l int) int {
	maxLen := 0
	position := 0
	mapTotal := make(map[int]int, l)
	for i := 0; i < l; i++ {
		if i > 0 {
			for j := i - 1; j >= 0; j-- {
				if arr[j].val != arr[i].val && canMove(arr[i], arr[j]) {
					continue
				}
			}
		}
		currentLen := calc(arr, i, l, 1, maxLen, mapTotal)
		if currentLen > maxLen {
			maxLen = currentLen
			position = i
		}
	}
	fmt.Println("position", position)
	return maxLen
}

func calc(arr []Item, current, l, currentLen, maxLen int, maxLenTotal map[int]int) int {
	if current > l-1 {
		return currentLen
	}
	// if (currentLen + l - current) <= maxLen {
	// 	return maxLen
	// }
	// fmt.Println("calc", current, l, currentLen, arr[current])
	potentials := []int{}
	k := current
	for k < l-1 && arr[k+1].val == arr[current].val {
		k++
	}
	for k < l {
		if canMove(arr[k], arr[current]) && arr[k].val != arr[current].val {
			potentials = append(potentials, k)
		}
		k++
	}
	if len(potentials) < 2 {
		for k+1 < l && arr[k].val == arr[k+1].val && canMove(arr[k], arr[k+1]) {
			potentials = append(potentials, current+1)
			k++
		}
	}
	fmt.Println("potentials", potentials)
	if len(potentials) == 0 {
		return currentLen
	}

	lx := 0
	lcc := 0
	for _, potential := range potentials {
		if lx >= (currentLen + l - 1 - potential) {
			continue
		}
		if total, ok := maxLenTotal[potential]; ok {
			lcc = total
		} else {
			lcc = calc(arr, potential, l, 1, maxLen, maxLenTotal)
			maxLenTotal[potential] = lcc
		}
		lx = max(lx, currentLen+lcc)
	}

	return lx
}

func canMove(pre, next Item) bool {
	return (pre.i == next.i) && math.Abs(float64(pre.j-next.j)) == 1 || (pre.j == next.j) && math.Abs(float64(pre.i-next.i)) == 1
}
