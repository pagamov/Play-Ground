package main

import (
	"fmt"
)

func maxProfit_int(prices []int) (int, int, int) {
	buy_i := -1
	sell_i := -1
	buy := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			buy_i = i
		} else if prices[i]-buy > profit {
			profit = prices[i] - buy
			sell_i = i
		}
	}
	fmt.Println(buy_i, sell_i)
	return profit, buy_i, sell_i
}

func maxProfit(prices []int) int {
	left := 0
	right := 0
	profit, buy_i, sell_i := maxProfit_int(prices)

	if buy_i > 0 {
		left = maxProfit(prices[:buy_i])
	}
	if sell_i < len(prices) && sell_i > 0 {
		right = maxProfit(prices[sell_i:])
	}
	return left + profit + right
}

func main() {
	test := [][]int{{7, 1, 5, 3, 6, 4}, {1, 2, 3, 4, 5}, {7, 6, 4, 3, 1}}
	for t := 0; t < len(test); t++ {
		fmt.Println(maxProfit(test[t]))
	}
}
