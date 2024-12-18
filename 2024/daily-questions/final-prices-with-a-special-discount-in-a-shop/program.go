package finalpriceswithaspecialdiscountinashop

type Element struct {
	Value int
	Index int
}

/*
prices
- price[i] i
discount: prices[j] <= prices[i] and j>i

Solution:
 - thinking about sort
 
 
*/

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}
