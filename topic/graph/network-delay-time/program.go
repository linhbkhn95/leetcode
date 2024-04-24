package networkdelaytime

type node struct {
}

func networkDelayTime(times [][]int, n int, k int) int {
	nodes := make(map[int][][]int, 0)
	for _, time := range times {
		nodes[time[0]] = append(nodes[time[0]], time)
	}

}

func graph(nodes map[int][][]int, n, source int, count *int, footprint map[int]bool, timeSignals map[int]int) {
	if footprint[source] {
		return
	}
	if *count == n {
		return
	}
	neighbors, ok := nodes[source]
	if !ok {
		return
	}
	for _, neighbor := range neighbors {
		totalTime, ok := timeSignals[neighbor[1]]
		if !ok {
			timeSignals[neighbor[1]] += neighbor[2]
		} else {
			if totalTime > timeSignals[source]+neighbor[2] {
				timeSignals[neighbor[1]] = timeSignals[source] + neighbor[2]
			}
		}
		graph(nodes, n, neighbor[1], count, footprint, timeSignals)
	}

}

// 1-2-3-4, 1-5-4
