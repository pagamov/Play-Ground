package main

func reverse(x int) int {
	res := 0
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	for x != 0 {
		pop := x % 10
		x /= 10
		res += pop
		res *= 10
	}
	res /= 10
	switch sign {
	case -1:
		res *= -1
	}

	if res > (1<<31)-1 || res < -1<<31 {
		return 0
	} else {
		return res
	}
}

func main() {

	reverse(123)

}
