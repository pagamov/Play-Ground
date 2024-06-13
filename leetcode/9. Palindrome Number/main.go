package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	start := x
	tmp := []int{}
	for x != 0 {
		tmp = append(tmp, x%10)
		x /= 10
	}
	sum := 0
	l := len(tmp)
	for i := 0; i < l; i++ {
		sum += tmp[i] * int(math.Pow(float64(10), float64(l-i-1)))
	}
	if start == sum {
		return true
	} else {
		return false
	}
}

func main() {
	test := []int{121, -12, 123321}

	for t := 0; t < len(test); t++ {
		fmt.Printf("%d - %t\n", test[t], isPalindrome(test[t]))
	}
}
