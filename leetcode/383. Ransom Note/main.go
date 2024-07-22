package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	mag := map[string]int{}
	for i := 0; i < len(magazine); i++ {
		mag[string(magazine[i])]++
	}
	for j := 0; j < len(ransomNote); j++ {
		if mag[string(ransomNote[j])] > 0 {
			mag[string(ransomNote[j])]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	test := [][]string{{"a", "b"}, {"aa", "ab"}, {"aa", "aab"}}
	for t := 0; t < len(test); t++ {
		fmt.Println(canConstruct(test[t][0], test[t][1]))
	}
}
