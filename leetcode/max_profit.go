package main

func maxProfit(prices []int) int {

	minPrice := int(^uint(0) >> 1)
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit

}

