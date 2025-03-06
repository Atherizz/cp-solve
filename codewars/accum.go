package main

import "strings"


func Accum(s string) string {
	res := ""
	for index, char := range s {
		res += strings.ToUpper(string(char))
		for i := 0; i < index; i++ {
			res += strings.ToLower(string(char))
		}
		if index != len(s)-1 {
			res += "-"
		}
	}
	return res
}