package lightbulbcircuit

type Node struct {
	IsTurnOn bool
	IsLight  bool
	Value    int32
}

func Solution(states []int32) (result int32) {
	result = 0
	circuitMapping := make(map[int32]*Node)

	for _, state := range states {
		node := &Node{
			Value:    state,
			IsTurnOn: true,
		}

		if state == 1 {
			node.IsLight = true
		}
		circuitMapping[state] = node
		if checkCircuitLight(circuitMapping[state], circuitMapping) {
			result++
		}
	}
	return
}

func checkCircuitLight(node *Node, circuitMapping map[int32]*Node) bool {
	if node.IsLight {
		return true
	}
	preNode, ok := circuitMapping[node.Value-1]
	if !ok {
		return false
	}
	if preNode.IsLight {
		return true
	}
	preNode.IsLight = checkCircuitLight(preNode, circuitMapping)
	return preNode.IsLight
}
