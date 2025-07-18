package main

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func ClosestPrimes(left int, right int) []int {
	primeNumber := []int{}

	for i := left; i <= right; i++ {
		if IsPrime(i) {
			primeNumber = append(primeNumber, i)
		}
	}

	if len(primeNumber) < 2 {
		return []int{-1, -1}
	}

    minGap := int(^uint(0) >> 1)
    res := []int{}

    for i := 0; i < len(primeNumber)-1; i++ {
        gap := primeNumber[i+1] - primeNumber[i]
        if gap < minGap {
            minGap = gap
            res = []int{primeNumber[i], primeNumber[i+1]}
        }
    }

	return res

}


