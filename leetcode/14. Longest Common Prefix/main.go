package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	// for i := 0; i < len(strs); i++ {
	// 	fmt.Printf("%s, ", strs[i])
	// }
	// fmt.Println()

	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
	}
	min_word := 201

	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < min_word {
			min_word = len(strs[i])
		}
	}

	fmt.Println(min_word)

	max_len := 0
	bigest_prefix := ""
	for max_len < min_word {
		letter := strs[0][max_len : max_len+1]
		for i := 0; i < len(strs); i++ {
			if strs[i][max_len:max_len+1] != letter {
				return bigest_prefix
			}
		}
		max_len += 1
		bigest_prefix += letter
	}
	return bigest_prefix
}

func main() {
	tests := [][]string{{"flower", "flow", "flight"}, {"dog", "racecar", "car"}, {"a"}}

	for t := 0; t < len(tests); t++ {
		fmt.Println(longestCommonPrefix(tests[t]))
	}
}
