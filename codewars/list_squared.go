package main

import (
	"math"
)

func ListSquared(m, n int) [][]int {
	var res [][]int

	for i := m; i < n; i++ {
		if Divisors(i) != nil {
			res = append(res, Divisors(i))
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res
}

func isPerfectSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt == math.Floor(sqrt)
}

func Divisors(n int) []int {
	var div []int
	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			div = append(div, (n/i)*(n/i))
			sum += (n / i) * (n / i)
		}
	}

	if isPerfectSquare(sum) {
		return []int{n, sum}
	}

	return nil
}
