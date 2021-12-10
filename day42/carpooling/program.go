package carpooling

import "sort"

func carPooling(trips [][]int, capacity int) bool {
	tripProcesssing := make(map[int]int, len(trips))
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})
	for _, trip := range trips {
		availableCapcity := ReturnPassenger(trip[1], tripProcesssing)
		capacity += availableCapcity
		if capacity < trip[0] {
			return false
		}
		_, ok := tripProcesssing[trip[2]]
		if !ok {
			tripProcesssing[trip[2]] = trip[0]
		} else {
			tripProcesssing[trip[2]] += trip[0]
		}
		capacity -= trip[0]
	}

	return true
}

func ReturnPassenger(start int, tripProcesssing map[int]int) int {
	count := 0
	for end, passengerCount := range tripProcesssing {
		if end <= start {
			count += passengerCount
			delete(tripProcesssing, end)
		}
	}
	return count
}
