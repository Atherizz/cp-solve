package main

import "fmt"

func closestPrimes(left int, right int) []int {
	res := []int{}
	var counter int
	primeNumber := []int{}


	for i := left; i <= right; i++ {
		counter = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				counter++
			}
		}
		if counter == 2 {
			primeNumber = append(primeNumber, i)
			// isPrime = append(isPrime, true)
		}
	}

	minGap := right - left

	for i := 0; i < len(primeNumber)-1; i++ {
		gap := primeNumber[i+1] - primeNumber[i]
		if gap < minGap {
			minGap = gap
			res = []int{primeNumber[i], primeNumber[i+1]}
		}
	}

	if len(primeNumber) < 2 {
		return []int{-1, -1}
	}

	return res

}

func main() {
	fmt.Println(closestPrimes(10, 19))
}
