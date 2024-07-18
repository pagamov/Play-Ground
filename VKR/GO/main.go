package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

var step int = 15000
var B int = 25

// var n int = 104729 * 21732382677641
var n int = 97 * 13

func eratosfen(n int) []int {
	var tmp []int = []int{}
	for i := 2; i < n+1; i++ {
		tmp = append(tmp, i)
	}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != 0 {
			var candidates []int = []int{}
			for j := 2 * tmp[i]; j < n+1; j += tmp[i] {
				candidates = append(candidates, j)
			}
			for j := 0; j < len(candidates); j++ {
				tmp[candidates[j]-2] = 0
			}
		}
	}
	var res []int = []int{-1}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != 0 {
			res = append(res, tmp[i])
		}
	}
	return res
}

func Q(x int, n int) int {
	var m int = int(math.Sqrt(float64(n))) + 1
	return (x+m)*(x+m) - n
}

func smooth(primes []int) map[int][]int {
	var smooth_numbers map[int][]int = make(map[int][]int)
	var pos int = 1
	for len(smooth_numbers) < B {
		var value = Q(pos, n)
		var coef []int = []int{}

		for i := 0; i < len(primes); i++ {
			coef = append(coef, 0)
		}

		if value < 0 {
			value *= -1
			coef[0] = 1
		}

		for i := 1; i < len(primes); i++ {
			for {
				if value%primes[i] == 0 {
					value /= primes[i]
					coef[i] += 1
				} else {
					break
				}
			}
		}

		if value == 1 {
			fmt.Println(pos, Q(pos, n), len(smooth_numbers), "/", B)
			smooth_numbers[pos] = coef
		}

		if pos > 0 {
			pos *= -1
		} else if pos < 0 {
			pos *= -1
			pos += 1
		}
	}
	return smooth_numbers
}

func main() {
	// fmt.Println("B " + color(string(rune(B)), "data"))

	start := time.Now()

	var primes []int = eratosfen(1000)

	// for i := 0; i < len(primes); i++ {
	// 	fmt.Println(primes[i])
	// }

	var smooth_numbers map[int][]int = smooth(primes)

	fmt.Println(len(smooth_numbers))

	var matrix [][]int = [][]int{}

	for 

	// res := factor(n, B)

	elapsed := time.Since(start)

	// fmt.Println(res)

	// print("\nans:", color(string(rune(res[0])), "strong")+" "+color(string(rune(res[1])), "strong"))

	log.Printf("Binomial took %s", elapsed)

}
