package main

// import "fmt"

func selectionSort(arr []int) []int {
	length := len(arr)
	temp := 0
	for i := 0; i < length - 1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		temp = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	return arr
}

func selectionSortDesc(arr []int) []int {
	length := len(arr)
	temp := 0
	for i := 0; i < length - 1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] > arr[minIndex] {
				minIndex = j
			}
		}
		temp = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	return arr
}

// func main() {
// 	fmt.Println(selectionSortDesc([]int{9, 1, 8, 2, 7, 3, 6, 4, 5}))
// }
