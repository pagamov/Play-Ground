package main

import "fmt"

func f(s string) {
	return ""
}

func main() {
	test := []string{}
	for t := 0; t < len(test); t++ {
		fmt.Println(f(test[t]))
	}
}
