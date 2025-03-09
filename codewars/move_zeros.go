package main

import "fmt"

func MoveZeros(arr []int) []int {
	var temp int

	for i := (len(arr) - 1); i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if arr[j] == 0 {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
}
