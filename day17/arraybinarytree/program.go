package arraybinarytree

func Solution(arr []int64) string {
	// Type your solution here
	l := len(arr)
	if l < 2 {
		return ""
	}
	countLeft := 1
	countRight := 1
	countL, countR := 0, 0
	index := 2
	exit := false

	calcCount := 0
	for index < l && calcCount < 1000000 {
		calcCount++

		cLeft := countLeft
		if cLeft > 0 {
			countL++
		}
		countLeft = 0
		for i := 0; i < cLeft; i++ {
			for j := 0; j < 2; j++ {
				if index+1 < l {
					if arr[index+1] != -1 {
						countLeft++
					}
					index++
				} else {
					exit = true
					break
				}
			}

		}
		if exit {
			break
		}
		cRight := countRight
		if cRight > 0 {
			countR++
		}
		countRight = 0
		for i := 0; i < cRight; i++ {
			for j := 0; j < 2; j++ {
				if index+1 < l {
					if arr[index+1] != -1 {
						countRight++
					}
					index++
				} else {
					exit = true
					break
				}
			}

		}
		if exit {
			break
		}
	}
	if countR > countL {
		return "Right"
	}
	if countR < countL {
		return "Left"
	}
	return ""
}
