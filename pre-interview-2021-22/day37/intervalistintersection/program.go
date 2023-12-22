package program

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := [][]int{}

	l1, l2 := len(firstList), len(secondList)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if firstList[i][0] <= secondList[j][0] && firstList[i][1] >= secondList[j][0] {
			start := secondList[j][0]
			end := firstList[i][1]
			if firstList[i][1] >= secondList[j][1] {
				end = secondList[j][1]
				j++
			} else {
				i++
			}
			result = append(result, []int{start, end})

		} else if secondList[j][0] <= firstList[i][0] && secondList[j][1] >= firstList[i][0] {
			start := firstList[i][0]
			end := secondList[j][1]
			if secondList[j][1] >= firstList[i][1] {
				end = firstList[i][1]
				i++
			} else {
				j++
			}
			result = append(result, []int{start, end})
		} else if firstList[i][0] > secondList[j][1] {
			j++
		} else {
			i++
		}
	}
	return result
}
