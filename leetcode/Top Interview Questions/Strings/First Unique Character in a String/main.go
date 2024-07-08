package main

import "fmt"

func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		_, exist := m[s[i]]
		if exist {
			m[s[i]] = -1
		} else {
			m[s[i]] = i
		}
	}
	min := 1000000
	for k, v := range m {
		if v != -1 && v < min {
			min = v
		}
		fmt.Println(string(k), v)
	}
	if min == 1000000 {
		return -1
	}
	return min
}

func main() {
	test := []string{"leetcode", "loveleetcode", "aabb"}
	for t := 0; t < len(test); t++ {
		fmt.Println(firstUniqChar(test[t]))
	}
}
