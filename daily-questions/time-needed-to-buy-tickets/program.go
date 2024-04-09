package timeneededtobuytickets

func timeRequiredToBuy(tickets []int, k int) int {
	result := 0
	for i, ticket := range tickets {
		if ticket < tickets[k] {
			result += ticket
		} else {
			if i <= k {
				result += tickets[k]
			} else {
				result += tickets[k] - 1
			}
		}
	}
	return result
}
