package main

func Sumation(n int) int {
	result := 0
	if n < 0 {
		return n
	} else {
		for i := 1; i <= n; i++ {
			result += i
		}
	}
	return result
}