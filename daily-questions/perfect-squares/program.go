package perfectsquares

import "math"

func numSquares(n int) int {
	if n < 4 {
		return n
	}

	sqrt := math.Floor(math.Sqrt(float64(n)))

}
