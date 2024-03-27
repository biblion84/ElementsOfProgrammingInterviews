package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(maxProfitTwoTrade([]int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}))

}

func maxProfitTwoTrade(prices []int) int {
	// bruteforce
	maxProfit := 0

	for i := range prices {
		firstTrade := maxProfitBook(prices[:i])
		secondTrade := maxProfitBook(prices[i:])

		maxProfit = max(maxProfit, firstTrade+secondTrade)
	}

	return maxProfit
}

func maxProfitBook(prices []int) int {
	minPrice := math.MaxInt
	profit := 0

	for _, price := range prices {
		profitDay := price - minPrice
		profit = max(profit, profitDay)
		minPrice = min(minPrice, price)
	}

	return profit
}
