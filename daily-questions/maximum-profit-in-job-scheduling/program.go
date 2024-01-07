package maximumprofitinjobscheduling

import (
	"sort"
)

type Job struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	footprint := make(map[int]int)
	jobs := make([]*Job, len(startTime))
	for i := range startTime {
		jobs[i] = &Job{
			startTime: startTime[i],
			endTime:   endTime[i],
			profit:    profit[i],
		}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].startTime < jobs[j].startTime
	})

	return dp(jobs, len(jobs), 0, footprint)
}

func dp(jobs []*Job, l, index int, footprint map[int]int) int {
	if index < 0 || index >= l {
		return 0
	}
	if v, ok := footprint[index]; ok {
		return v

	}
	profit1 := dp(jobs, l, index+1, footprint)
	nextIndex := sort.Search(l, func(i int) bool {
		return jobs[i].startTime >= jobs[index].endTime
	})
	profit2 := jobs[index].profit + dp(jobs, l, nextIndex, footprint)
	profit := max(profit1, profit2)
	footprint[index] = profit
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
