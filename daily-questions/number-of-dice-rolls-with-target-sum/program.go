package numberofdicerollswithtargetsum

import (
	"fmt"
	"math"
)

var tenPowNine = int64(math.Pow(10, 9))

func numRollsToTarget(n int, k int, target int) int {
	footprint := make(map[string]int64)
	return int(tryRolling(n, 0, k, 0, target, footprint))
}

func tryRolling(n, stt, k, current, target int, footprint map[string]int64) int64 {
	if stt == n {
		if current == target {
			return 1
		}
		return 0
	}

	var total int64
	for i := 1; i <= k; i++ {
		count, ok := footprint[getKey(stt+1, k, current+i)]
		if ok {
			total += count
			continue
		}
		total += tryRolling(n, stt+1, k, current+i, target, footprint)
	}
	total = total % (tenPowNine + 7)
	footprint[getKey(stt, k, current)] = total
	return int64(total)
}

func getKey(stt, k, current int) string {
	return fmt.Sprintf("%d_%d_%d", stt, k, current)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
