package main

import (
	"fmt"
)

// func maxProfit(prices []int) int {
// 	max := 0
// 	for i := 0; i < len(prices)-1; i++ {
// 		if prices[i] > prices[i+1] {
// 			continue
// 		}
// 		for j := i + 1; j < len(prices); j++ {
// 			dif := prices[j] - prices[i]
// 			if dif > max {
// 				max = dif
// 			}
// 		}
// 	}
// 	return max
// }

func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else if prices[i]-buy > profit {
			profit = prices[i] - buy
		}
	}
	return profit
}

func main() {
	test := [][]int{{7, 1, 5, 3, 6, 4}, {7, 6, 4, 3, 1}}
	for t := 0; t < len(test); t++ {
		fmt.Println(maxProfit(test[t]))
	}
}
