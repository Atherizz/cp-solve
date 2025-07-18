package main

import "fmt"

func divide(dividend int, divisor int) int {

	res := dividend
	if dividend == 0 || divisor == 0 {
		return 0
	}

	i := 0
	for res > 0 {
		res -= abs(divisor)
		i++
	}

	result := i - 1

	if dividend == divisor {
		result = dividend
	}

	if divisor < 0 {
		result *= -1
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(divide(10, 3))
}
