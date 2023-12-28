package countzerorequestservers

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	indexedQueries := make([][3]int, len(queries))
	for i := 0; i < len(queries); i++ {
		indexedQueries[i] = [3]int{i, queries[i] - x, queries[i]}
	}
	sort.Slice(indexedQueries, func(i, j int) bool {
		return indexedQueries[i][1] < indexedQueries[j][1]
	})
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1] // sort by time
	})
	// Now let's proceed processing queries sorted by time. As we go further
	// we'll add all active servers to a map/set. We'll ensure that any server
	// outside the window is removed
	active := make(map[int]int)
	arr := make([]int, len(queries))
	left, right := 0, 0
	for i := 0; i < len(indexedQueries); i++ {
		idx, start, end := indexedQueries[i][0], indexedQueries[i][1], indexedQueries[i][2]
		// Now let's add all servers that are active
		for right < len(logs) && logs[right][1] <= end {
			active[logs[right][0]]++
			right++
		}
		// First let's remove any server that is outside the window
		for left < len(logs) && logs[left][1] < start {
			active[logs[left][0]]--
			if active[logs[left][0]] == 0 {
				delete(active, logs[left][0])
			}
			left++
		}
		arr[idx] = n - len(active)
	}
	return arr
}
func countServers1(n int, logs [][]int, x int, queries []int) []int {
	result := make([]int, len(queries))
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	sort.Slice(queries, func(i, j int) bool {
		return queries[i] < queries[j]
	})

	for i, query := range queries {
		result[i] = countNumberOfIdleServers(n, logs, query-x, query)
	}
	return result
}

func countNumberOfIdleServers(n int, logs [][]int, startTime, endTime int) int {
	appearedServer := make(map[int]struct{}, n)
	if logs[len(logs)-1][1] < startTime || logs[0][1] > endTime {
		return n
	}
	for _, log := range logs {
		if _, ok := appearedServer[log[0]]; ok {
			continue
		}
		if startTime <= log[1] && log[1] <= endTime {
			appearedServer[log[0]] = struct{}{}
		}
	}
	return n - len(appearedServer)
}

// [1,3] [1,5] [2,6]

// [5,10] [7,12]
//
// 1 2 3 4 5 6 7
