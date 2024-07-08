package main

import "fmt"

func myAtoi(s string) int {
	var res int = 0

	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i])
	}

	return res
}

func main() {
	test := []string{"42", " -042"}
	for t := 0; t < len(test); t++ {
		fmt.Println(myAtoi(test[t]))
	}
}
