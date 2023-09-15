package besttimetobuyandsellstocklc121

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for currDay := 1; currDay < len(prices); currDay++ {
		if prices[currDay] > minPrice {
			if profit := prices[currDay] - minPrice; profit > maxProfit {
				maxProfit = profit
			}
		} else {
			minPrice = prices[currDay]
		}
	}
	return maxProfit
}
