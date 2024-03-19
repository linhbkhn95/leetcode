package task_scheduler

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	cnt := make([]int, 26)
	for _, task := range tasks {
		cnt[task-'A']++
	}

	var maxCount, sameMaxCount int
	for _, c := range cnt {
		if c > maxCount {
			maxCount = c
			sameMaxCount = 1
		} else if c == maxCount {
			sameMaxCount++
		}
	}

	res := (n+1)*(maxCount-1) + sameMaxCount
	if res > len(tasks) {
		return res
	} else {
		return len(tasks)
	}
}

//type TaskFre struct {
//	T   byte
//	Fre int
//}
//
//func leastInterval(tasks []byte, n int) int {
//	taskFre := make(map[byte]int)
//	for _, t := range tasks {
//		taskFre[t]++
//	}
//	tasksFres := make([]*TaskFre, 0, len(taskFre))
//	for t, fre := range taskFre {
//		tasksFres = append(tasksFres, &TaskFre{
//			T:   t,
//			Fre: fre,
//		})
//	}
//	sort.Slice(tasksFres, func(i, j int) bool {
//		return tasksFres[i].Fre < tasksFres[j].Fre
//	})
//	i, j := 0, 0
//	result := 0
//	distance := 0
//	for i < len(tasksFres) {
//		if tasksFres[i].Fre == 0 {
//			i++
//			continue
//		}
//		if i < len(tasksFres) && j == len(tasksFres) {
//			result += n - distance + 1
//		}
//		distance = 0
//		j = i
//		for j < len(tasksFres) && distance <= n {
//			result++
//			distance++
//			tasksFres[j].Fre--
//			j++
//		}
//
//	}
//	return result
//}
