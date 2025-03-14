package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		middle := (low+high)/2 
		if target == arr[middle] {
			return middle
		} else if target > arr[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{3, 7, 12, 18, 22, 27, 31, 35, 39, 43, 47, 52, 58, 63, 68, 73, 79, 84, 91, 98}, 35))
}
