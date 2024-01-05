package imagesmoother

import "math"

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = img[i][j]
		}
	}
	i, j := 0, 0
	for i < m && i+2 < m {
		mStart := i
		mEnd := i + 2
		for j < n && j+2 < n {
			nStart := j
			nEnd := j + 2
			roundingDown(img, result, mStart, mEnd, nStart, nEnd)
			j += 2
		}
		i += 3
	}
	return result
}

func roundingDown(img, result [][]int, mStart, mEnd, nStart, nEnd int) {
	if nStart == 0 {
		result[mStart][nStart] = int(math.Floor(float64((img[mStart][nStart] + img[mStart][nStart+1] + img[mStart+1][nStart] + img[mStart+1][nStart+1]) / 4)))
		result[mStart+1][nStart] = int(math.Floor(float64((img[mStart+1][nStart] + img[mStart+1][nStart+1] + img[mStart][nStart+1] + img[mStart+2][nStart] + img[mStart][nStart] + img[mStart+2][nStart+1]) / 6)))
		result[mStart+2][nStart] = int(math.Floor(float64((img[mStart+1][nStart] + img[mStart+2][nStart] + img[mStart+2][nStart+1] + img[mStart+1][nStart+1]) / 4)))
	}

	result[mStart][nStart+1] = int(math.Floor(float64((img[mStart][nStart+1] + img[mStart][nStart] + img[mStart+1][nStart+1] + img[mStart][nStart+2] + img[mStart+1][nStart] + img[mStart+1][nStart+2]) / 6)))

	result[mStart][nStart+1] = int(math.Floor(float64((img[mStart][nStart+1] + img[mStart][nStart] + img[mStart+1][nStart+1] + img[mStart][nStart+2] + img[mStart+1][nStart] + img[mStart+1][nStart+2]) / 6)))
	result[mStart][nStart+2] = int(math.Floor(float64((img[mStart][nStart+2] + img[mStart][nStart+1] + img[mStart+1][nStart+1] + img[mStart+1][nStart+2]) / 4)))

	result[mStart+1][nStart+1] = int(math.Floor(float64((img[mStart][nStart] + img[mStart][nStart+1] + img[mStart][nStart+2] + img[mStart+1][nStart] + img[mStart+1][nStart+1] + img[mStart+1][nStart+2] + img[mStart+2][nStart] + img[mStart+2][nStart+1] + img[mStart+2][nStart+2]) / 9)))
	result[mStart+1][nStart+2] = int(math.Floor(float64((img[mStart+1][nStart+2] + img[mStart+1][nStart+1] + img[mStart][nStart+2] + img[mStart+2][nStart+2] + img[mStart][nStart+1] + img[mStart+2][nStart+1]) / 6)))

	result[mStart+2][nStart+1] = int(math.Floor(float64((img[mStart+1][nStart] + img[mStart+1][nStart+1] + img[mStart+1][nStart+2] + img[mStart+2][nStart] + img[mStart+2][nStart+2] + img[mStart+2][nStart+1]) / 6)))
	result[mStart+2][nStart+2] = int(math.Floor(float64((img[mStart+1][nStart+2] + img[mStart+1][nStart+1] + img[mStart+2][nStart+2] + img[mStart+2][nStart+1]) / 4)))
}
