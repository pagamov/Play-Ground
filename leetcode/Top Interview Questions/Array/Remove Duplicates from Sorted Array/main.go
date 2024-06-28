package main

import "fmt"

func removeDuplicates(nums []int) int {
	dub := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dub += 1
		}
		nums[i-dub] = nums[i]
	}
	for i := 0; i < dub; i++ {
		fmt.Printf("%d,", nums[i])
	}
	return len(nums) - dub
}

func main() {
	test := [][]int{{1, 1, 2}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}
	for t := 0; t < len(test); t++ {
		fmt.Println(removeDuplicates(test[t]))
	}
}
