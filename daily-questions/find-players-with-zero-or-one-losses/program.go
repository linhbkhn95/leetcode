package main

import (
	"slices"
)

func findWinners(matches [][]int) [][]int {
	userStats := make(map[int]int, 0)
	for _, match := range matches {
		player1Idx, player2Idx := match[0], match[1]
		if _, ok := userStats[player1Idx]; !ok {
			userStats[player1Idx] = 0
		}
		userStats[player2Idx]++
	}
	result := make([][]int, 2)
	for user, num := range userStats {
		if num == 0 {
			result[0] = append(result[0], user)
		}
		if num == 1 {
			result[1] = append(result[1], user)
		}
	}
	if result[0] != nil {
		slices.Sort(result[0])
	}
	if result[1] != nil {
		slices.Sort(result[1])
	}
	return result
}
