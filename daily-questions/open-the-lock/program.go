package openthelock

import "container/list"

type pair struct {
	node string
	step int
}

func neighbors(node string) []string {
	result := []string{}
	for i := range node {
		for _, d := range []int{1, -1} {
			neighbor := []byte(node)
			move := int(neighbor[i]) + d
			switch {
			case move < 48:
				move = 57
			case move > 57:
				move = 48
			}
			neighbor[i] = byte(move)
			result = append(result, string(neighbor))
		}
	}

	return result
}

func openLock(deadends []string, target string) int {
	seen := make(map[string]bool)
	for _, deadend := range deadends {
		if deadend == "0000" {
			return -1
		}
		seen[deadend] = true
	}
	seen["0000"] = true

	queue := list.New()
	queue.PushBack(pair{"0000", 0})

	for queue.Len() > 0 {
		current := queue.Front().Value.(pair)
		queue.Remove(queue.Front())

		if current.node == target {
			return current.step
		}

		for _, neighbor := range neighbors(current.node) {
			if !seen[neighbor] {
				seen[neighbor] = true
				queue.PushBack(pair{neighbor, current.step + 1})
			}
		}
	}

	return -1
}
