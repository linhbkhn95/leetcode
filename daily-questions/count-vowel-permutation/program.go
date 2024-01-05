package countvowelpermutation

import "math"

// func countVowelPermutation(n int) int {

// }

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs := 0
	testRound := minutesToTest/minutesToDie + 1
	for math.Pow(float64(testRound), float64(pigs)) <= float64(buckets) {
		pigs++
	}
	return pigs
}
