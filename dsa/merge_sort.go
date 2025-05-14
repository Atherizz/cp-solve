package main

// import "fmt"

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	leftArr := MergeSort(arr[:mid])
	rightArr := MergeSort(arr[mid:])

	Merge(leftArr, rightArr, arr)

	return arr
}

func Merge(leftArr []int, rightArr []int, arr []int) {
	
}

// func main() {
// 	fmt.Println(MergeSort([]int{9, 1, 8, 2, 7, 3, 6, 4, 5}))
// }