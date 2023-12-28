package singlethreadedcpu

import (
	"container/heap"
	"sort"
)

func getOrder(tasks [][]int) []int {
	l := len(tasks)
	availableTasks := make([]*AvaibleTask, len(tasks))
	for i, task := range tasks {
		availableTasks[i] = &AvaibleTask{
			Index:          i,
			ProcessingTime: task[1],
			EnqueueTime:    task[0],
		}
	}
	sort.Slice(availableTasks, func(i, j int) bool {
		if availableTasks[i].EnqueueTime == availableTasks[j].EnqueueTime {
			return availableTasks[i].ProcessingTime < availableTasks[j].ProcessingTime
		}
		return availableTasks[i].EnqueueTime < availableTasks[j].EnqueueTime
	})
	result := make([]int, 0, len(tasks))
	result = append(result, availableTasks[0].Index)
	endTime := availableTasks[0].ProcessingTime + availableTasks[0].EnqueueTime
	i := 1
	taskQueue := &AvailableTaskQueue{}
	for {
		if i == l {
			break
		}
		for i < l && availableTasks[i].EnqueueTime <= endTime {
			heap.Push(taskQueue, availableTasks[i])
			i++
		}
		if taskQueue.Len() > 0 {
			candicate := heap.Pop(taskQueue)
			availTask := candicate.(*AvaibleTask)
			result = append(result, availTask.Index)
			endTime = endTime + availTask.ProcessingTime
		} else if i < l {
			result = append(result, availableTasks[i].Index)
			endTime = availableTasks[i].EnqueueTime + availableTasks[i].ProcessingTime
			i++
		}
	}
	for taskQueue.Len() > 0 {

		candicate := heap.Pop(taskQueue)
		availTask := candicate.(*AvaibleTask)
		result = append(result, availTask.Index)
	}
	return result
}

type AvaibleTask struct {
	Index          int
	ProcessingTime int
	EnqueueTime    int
}

type AvailableTaskQueue []*AvaibleTask

// Len implements heap.Interface.
func (a AvailableTaskQueue) Len() int {
	return len(a)
}

// Less implements heap.Interface.
func (a AvailableTaskQueue) Less(i int, j int) bool {
	if a[i].ProcessingTime == a[j].ProcessingTime {
		return a[i].Index < a[j].Index
	}
	return a[i].ProcessingTime < a[j].ProcessingTime
}

// Pop implements heap.Interface.
func (a *AvailableTaskQueue) Pop() any {

	item := (*a)[a.Len()-1]
	(*a) = (*a)[:a.Len()-1]
	return item
}

// Push implements heap.Interface.
func (a *AvailableTaskQueue) Push(x any) {
	*a = append(*a, x.(*AvaibleTask))
}

// Swap implements heap.Interface.
func (a AvailableTaskQueue) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

var _ heap.Interface = (*AvailableTaskQueue)(nil)
