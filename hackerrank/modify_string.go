package main

import "fmt"

func ModifyString(str string) string {
	res := ""
	length := len(str) - 1
	for i := length; i >= 0; i-- {
		if str[i] < '1' || str[i] > '9' {
			res += string(str[i])
		}
	}

	return res
}

func main() {
	fmt.Println(ModifyString("de75s1rev2er"))
}
