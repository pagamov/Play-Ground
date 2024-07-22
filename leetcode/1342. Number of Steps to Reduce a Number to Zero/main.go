package main

import "fmt"

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	} else if num%2 == 0 {
		return numberOfSteps(num/2) + 1
	} else {
		return numberOfSteps(num-1) + 1
	}
}

func main() {
	test := []int{14, 8, 123}
	for t := 0; t < len(test); t++ {
		fmt.Println(numberOfSteps(test[t]))
	}
}
