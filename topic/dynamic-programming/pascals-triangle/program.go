package pascalstriangle

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	if numRows == 1 {
		return result
	}
	result[1] = []int{1, 1}
	if numRows == 2 {
		return result
	}
	for i := 2; i < numRows; i++ {
		row := []int{1}
		for j := 0; j < len(result[i-1])-1; j++ {
			row = append(row, result[i-1][j]+result[i-1][j+1])
		}
		row = append(row, 1)
		result[i] = row
	}
	return result
}
