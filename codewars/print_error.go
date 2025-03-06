package main

import "strconv"

func PrinterError(s string) string {
	counter := 0
	res := ""
	for _, char := range s {
		if char < 'a' || char > 'm' {
			counter++
		}
	}

	res += strconv.Itoa(counter) + "/" + strconv.Itoa(len(s))
	return res
}