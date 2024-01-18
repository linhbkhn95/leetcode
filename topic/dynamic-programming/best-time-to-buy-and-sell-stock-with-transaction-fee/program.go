package besttimetobuyandsellstockwithtransactionfee

func maxProfit(prices []int, fee int) int {
	footprint := make(map[key]int)
	return dp(prices, len(prices), -1, 0, fee, footprint)
}

type key struct {
	lastBuyIndex int
	index        int
}

func dp(prices []int, l, lastBuyIndex, index int, fee int, footprint map[key]int) int {
	if index == l {
		return 0
	}
	k := key{
		lastBuyIndex: lastBuyIndex,
		index:        index,
	}
	if v, ok := footprint[k]; ok {
		return v
	}

	var maxProfit int
	if lastBuyIndex != -1 {
		maxProfit = max(prices[index]-prices[lastBuyIndex]-fee+dp(prices, l, -1, index+1, fee, footprint), dp(prices, l, lastBuyIndex, index+1, fee, footprint))
	} else {
		maxProfit = max(dp(prices, l, -1, index+1, fee, footprint), dp(prices, l, index, index+1, fee, footprint))
	}

	footprint[k] = maxProfit
	return maxProfit
}
