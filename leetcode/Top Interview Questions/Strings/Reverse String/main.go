package main

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	test := [][]byte{{'h', 'e', 'l', 'l', 'o'}, {'H', 'a', 'n', 'n', 'a', 'h'}}
	for t := 0; t < len(test); t++ {
		reverseString(test[t])
	}
}
