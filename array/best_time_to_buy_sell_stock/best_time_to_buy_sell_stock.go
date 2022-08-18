package main

func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]

	for _, price := range prices {
		if price < buy {
			buy = price
			continue
		}
		if profit < price-buy {
			profit = price - buy
		}
	}
	return profit
}
