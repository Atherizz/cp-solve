package main

import (
	"strconv"
	"strings"
)

func Order(sentence string) string {
	separated := strings.Split(sentence, " ")
	order := make([]string, len(separated))

	for _, word := range separated {
		for _, char := range word {
			if char >= '0' && char <= '9' {
				num, _ := strconv.Atoi(string(char))
				order[num-1] = word
			}
		}
	}

	return strings.Join(order, " ")
}
