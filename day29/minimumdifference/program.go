package program

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	l := len(arr)
	if l == 2 {
		return [][]int{{arr[0], arr[1]}}
	}
	sort.Ints(arr)
	var result [][]int

	minDistance := arr[1] - arr[0]
	for i := 1; i < l; i++ {
		distance := arr[i] - arr[i-1]
		if minDistance > distance {
			minDistance = distance
			result = [][]int{
				{
					arr[i-1], arr[i],
				},
			}
		} else if minDistance == distance {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
