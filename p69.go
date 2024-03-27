package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(maxProfit([]int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}))
	fmt.Println(maxProfitBook([]int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}))

}

func maxProfit(in []int) int {
	profit := 0

	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			profit = max(in[j]-in[i], profit)
		}
	}

	return profit
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
