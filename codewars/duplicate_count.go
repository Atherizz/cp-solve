package main

import "strings"

func duplicate_count(s1 string) int {
	counter := 0
	data := make(map[string]int)

	word := strings.ToLower(s1)

	for _, letter := range word {
		char := string(letter)
		data[char]++
	}

	for _, value := range data {
		if value > 1 {
			counter++
		}
	}
	return counter

}