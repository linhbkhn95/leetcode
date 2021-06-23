package schedulerttask

import (
	"fmt"
	"sort"
)

type taskValue struct {
	taskCount   int
	latestIndex int
}

type machineTask map[byte]*taskValue

func leastInterval(tasks []byte, n int) int {

	machine := make(machineTask)
	var tks []byte
	for _, task := range tasks {
		tv, ok := machine[task]
		if ok {
			machine[task].taskCount = tv.taskCount + 1
		} else {
			tks = append(tks, task)
			machine[task] = &taskValue{
				taskCount:   1,
				latestIndex: -1,
			}
		}
	}

	sort.Slice(tks, func(i, j int) bool {
		return machine[tks[i]].taskCount > machine[tks[j]].taskCount
	})
	index := 0
	stt := 0
	var works []string
	count := 0
	loop := 0
	fmt.Println("tks", tks)

	for len(tks) > 0 && loop < 20 {
		loop++
		if index >= len(tks) {
			index = 0
		}
		fmt.Println("adasdasdasd", tks, count, index)

		// fmt.Println("tks", tks, index)
		task := tks[index]
		value := machine[task]
		if value == nil {
			continue
		}
		fmt.Println("adasdasdasd", task, value, count, index)

		if value.latestIndex == -1 || count == n || stt-value.latestIndex >= n {

			if value.taskCount == 1 {
				tks = remove(tks, index)
			}
			machine.doTask(task, value.taskCount, stt)

			works = append(works, string(task))
			stt += 1
			if len(tks) == 0 {
				break
			}
			if count == n {
				index = 0
				count = 0
			} else {
				index++
				count++
			}
		} else {
			stt += 1
			count++
			works = append(works, "rest")
		}

	}
	fmt.Println("works", works)
	return stt
}

func remove(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}

func (m machineTask) doTask(task byte, taskCount, index int) {
	if taskCount == 1 {
		delete(m, task)
		return
	}
	m[task].latestIndex = index
	m[task].taskCount = taskCount - 1
}
