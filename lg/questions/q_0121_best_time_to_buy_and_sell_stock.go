package questions

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	i, j := 0, 1
	maxProfit := 0
	for j < len(prices) {
		currentMaxProfit := prices[j] - prices[i]
		maxProfit = max(currentMaxProfit, maxProfit)
		if prices[i] >= prices[j] {
			i = j
			j++
		} else {
			j++
		}
	}

	return maxProfit
}
