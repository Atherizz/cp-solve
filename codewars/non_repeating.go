package main

import (
	"strings"
)

func FirstNonRepeating(str string) string {
    data := make(map[rune]int)

    for _, letter := range str {
        lowerChar := rune(strings.ToLower(string(letter))[0])
        data[lowerChar]++
    }

    for _, letter := range str {
        lowerChar := rune(strings.ToLower(string(letter))[0])
        if data[lowerChar] == 1 {
            return string(letter)
        }
    }

    return ""
}

