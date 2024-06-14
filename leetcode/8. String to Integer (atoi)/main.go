package main

import "fmt"

func myAtoi(s string) int {
	fmt.Printf("test - %s\n", s)
	number := []int{}
	sign := 1
	start := false
	for i := 0; i < len(s); i++ {
		switch s[i : i+1] {
		case "-":
			if !start {
				sign = -1
			}
		case "0":
			if start {
				number = append(number, 0)
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if !start {
				start = true
				number = append(number, int(s[i : i+1][0]-'0'))
			}
		default:
			if start {

			}
		}
	}
	return 0
}

func main() {
	test := []string{"42", " -042", "1337c0d3", "0-1", "words and 987"}

	for t := 0; t < len(test); t++ {
		fmt.Println(myAtoi(test[t]))
	}
}
