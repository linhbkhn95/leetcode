package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	nodes := make(map[int][]int)
	indexes := make(map[string]int)
	for _, flight := range flights {
		from, to, cost := flight[0], flight[1], flight[2]
		nodes[from] = append(nodes[from], to)
		str := fmt.Sprintf("%d:%d", from, to)
		indexes[str] = cost
	}
	result := math.MaxInt32

	walk(nodes, indexes, k+1, src, dst, 0, &result, map[int]bool{})
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

func walk(nodes map[int][]int, indexes map[string]int, k, node, dst, currentPrice int, cheapestPrice *int, footprint map[int]bool) {
	if k < 0 {
		return
	}
	if dst == node {
		*cheapestPrice = min(*cheapestPrice, currentPrice)
		return
	}
	for _, to := range nodes[node] {
		if footprint[to] {
			continue
		}
		str := fmt.Sprintf("%d:%d", node, to)

		footprint[to] = true
		walk(nodes, indexes, k-1, to, dst, currentPrice+indexes[str], cheapestPrice, footprint)
		delete(footprint, to)
	}
}
