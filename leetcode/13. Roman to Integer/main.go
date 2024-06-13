package main

import "fmt"

func romanToInt(s string) int {
	sym := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	tmp := 0
	leng := len(s)
	for i := 0; i < leng; i++ {
		if i < leng-1 {
			switch s[i : i+2] {
			case "IV":
				tmp += 4
				i++
			case "IX":
				tmp += 9
				i++
			case "XL":
				tmp += 40
				i++
			case "XC":
				tmp += 90
				i++
			case "CD":
				tmp += 400
				i++
			case "CM":
				tmp += 900
				i++
			default:
				tmp += sym[s[i:i+1]]
			}
		} else {
			tmp += sym[s[i:i+1]]
		}
	}
	return tmp
}

func main() {
	t := []string{"III", "LVIII", "MCMXCIV"}

	for test := 0; test < len(t); test++ {
		fmt.Println(t[test])
		fmt.Println(romanToInt(t[test]))
		fmt.Println()
	}
}
