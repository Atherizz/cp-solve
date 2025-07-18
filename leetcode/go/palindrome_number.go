package main

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	index := len(s) - 1

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[index] {
			return false
		}
		index--
	}
	return true
}


