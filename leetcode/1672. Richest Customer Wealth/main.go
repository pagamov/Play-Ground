package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	max := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	test := [][][]int{{{1, 2, 3}, {3, 2, 1}}, {{1, 5}, {7, 3}, {3, 5}}, {{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}}
	for t := 0; t < len(test); t++ {
		fmt.Println(maximumWealth(test[t]))
	}
}
