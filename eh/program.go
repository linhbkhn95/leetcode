package eh

import (
	"strings"
)

func Solution1(x string, y string) bool {
	l1, l2 := len(x), len(y)
	if l1 < l2 {
		return false
	}
	i, j := 0, 0
	count := 0
	for i < l1 && j < l2 {
		if x[i] == y[j] {
			count++
			j++
		}
		i++
	}
	return count == l2
}

func Solution(connections []string, name1 string, name2 string) int {
	// Implement your solution here
	mapping := make(map[string][]string, 0)
	for _, conn := range connections {
		names := strings.Split(conn, ":")
		mapping[names[0]] = append(mapping[names[0]], names[1])
		mapping[names[1]] = append(mapping[names[1]], names[0])
	}
	minCount := len(connections)
	footprint := make(map[string]struct{}, 0)
	dp(mapping, footprint, name1, name2, &minCount, 0)

	if minCount == len(connections) {
		return -1
	}
	return minCount
}

func dp(mapping map[string][]string, footprint map[string]struct{}, name1, name2 string, minCount *int, currentCount int) {

	if name1 == name2 {
		*minCount = min(*minCount, currentCount)
		return
	}
	friends, ok := mapping[name1]
	if !ok {
		return
	}
	for _, friend := range friends {
		if _, ok := footprint[friend]; ok {
			continue
		}
		footprint[friend] = struct{}{}
		dp(mapping, footprint, friend, name2, minCount, currentCount+1)
		delete(footprint, friend)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
