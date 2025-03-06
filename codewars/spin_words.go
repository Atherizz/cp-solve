package main

import "strings"

func SpinWords(str string) string {
	separated := strings.Split(str, " ")
	res := ""

	for i, word := range separated {
		if len(word) >= 5 {
			for i := (len(word) - 1); i >= 0; i-- {
				char := string(word[i])
				res += char
			}
		} else {
			for _, char := range word {
				letter := string(char)
				res += letter
			}
		}
		if i != len(separated)-1 {
			res += " "
		}
	}
	return res
}