package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	} else {
		prevRow := getRow(rowIndex - 1)
		row := make([]int, rowIndex+1)
		row[0] = 1
		for i := 1; i < rowIndex; i++ {
			row[i] = prevRow[i-1] + prevRow[i]
		}
		row[rowIndex] = 1
		return row
	}
}

func main() {
	test := []int{3, 0, 1}
	for t := 0; t < len(test); t++ {
		fmt.Println(getRow(test[t]))
	}
}
