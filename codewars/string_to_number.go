package main

import "strconv"

func StringToNumber(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}