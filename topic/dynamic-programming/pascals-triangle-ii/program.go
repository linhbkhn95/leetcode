package pascalstriangleii

func getRow(rowIndex int) []int {
	pascalTriangle := make([][]int, 0)
	pascalTriangle = append(pascalTriangle, []int{1})
	pascalTriangle = append(pascalTriangle, []int{1, 1})
	for i := 2; i <= rowIndex; i++ {
		row := make([]int, 0)
		row = append(row, 1)
		for j := 0; j < len(pascalTriangle[i-1])-1; j++ {
			row = append(row, pascalTriangle[i-1][j]+pascalTriangle[i-1][j+1])
		}
		row = append(row, 1)
		pascalTriangle = append(pascalTriangle, row)
	}
	return pascalTriangle[rowIndex]
}
