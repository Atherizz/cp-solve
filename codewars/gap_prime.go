package main

import (
	"math"
)

func Gap(g, m, n int) []int {
	res := PrimeNumber(m, n)

	for i := 0; i < len(res)-1; i++ {
		if res[i+1]-res[i] == g {
			return []int{res[i], res[i+1]}
		}
	}

	return nil
}

func PrimeNumber(m, n int) []int {
	var prime []int
	for i := m; i <= n; i++ {
		if IsPrime(i) {
			prime = append(prime, i)
		}
	}
	
	return prime
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}

