package main

import "fmt"

func fizzBuzz(n int) []string {
	res := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	return res
}

func main() {
	test := []int{3, 5, 15}
	for t := 0; t < len(test); t++ {
		fmt.Println(fizzBuzz(test[t]))
	}
}
