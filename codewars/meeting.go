package main

import (
	"strings"
)

func Meeting(s string) string {
	res := ""
	uppercase := strings.ToUpper(s)
	names := strings.Split(uppercase, ";")
	firstName := make([]string, len(names))
	lastName := make([]string, len(names))

	for i, name := range names {
		splitName := strings.Split(name, ":")
		firstName[i] = splitName[0]
		lastName[i] = splitName[1]
	}

	// sort last name
	for i := 0; i < len(lastName)-1; i++ {
		for j := 0; j < len(lastName)-i-1; j++ {
			if lastName[j] > lastName[j+1] || lastName[j] == lastName[j+1] && firstName[j] > firstName[j+1] {
				temp := lastName[j]
				lastName[j] = lastName[j+1]
				lastName[j+1] = temp

				temp2 := firstName[j]
				firstName[j] = firstName[j+1]
				firstName[j+1] = temp2
			}
		}
	}

	for i, _ := range names {
		res += "(" + lastName[i] + "," + firstName[i] + ")"
	}

	return res

}

