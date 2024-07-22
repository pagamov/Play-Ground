package main

import (
	"fmt"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	m := s
	if strings.ToLower(clearString(s)) == Reverse(strings.ToLower(clearString(m))) {
		return true
	}
	return false
}

func main() {
	test := []string{"A man, a plan, a canal: Panama", "race a car"}
	for t := 0; t < len(test); t++ {
		fmt.Println(isPalindrome(test[t]))
	}
}
