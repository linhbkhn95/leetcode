package timeneededtoinformallemployees

type Node struct {
	Val        int
	InformTime int
	Chirls     []*Node
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	nodes := make(map[int]*Node)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{
			Val:        manager[i],
			InformTime: informTime[i],
		}
	}

	var root *Node
	for i := 0; i < n; i++ {
		parent, ok := nodes[manager[i]]
		if !ok {
			root = nodes[i]
			continue
		}
		parent.Chirls = append(parent.Chirls, nodes[i])
	}

	return bfs(root)
}

func bfs(root *Node) int {
	if root == nil {
		return 0
	}

	maxV := 0
	for _, child := range root.Chirls {
		maxV = max(maxV, root.InformTime+bfs(child))
	}
	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
