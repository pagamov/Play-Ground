package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("B " + color(string(B), "data"))

	start := time.Now()

	res = Factor(n, B)

	elapsed := time.Since(start)

	fmt.Println(res)

	print("\nans:", color(int(res[0]), "strong")+" "+color(int(res[1]), "strong"))

	log.Printf("Binomial took %s", elapsed)

}
