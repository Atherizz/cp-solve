package main

import "strconv"

func CreatePhoneNumber(numbers [10]uint) string {
	result := "("
	for i := 0; i < 3; i++ {
		result += strconv.Itoa(int(numbers[i]))
	}
	result += ") "
	for i := 3; i < len(numbers); i++ {
		if i == 6 {
			result += "-"
		}
		result += strconv.Itoa(int(numbers[i]))
	}

	return result
}