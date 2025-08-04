package main

import "fmt"

func CreateSpiral(n int) [][]int {

	if n <= 0 {
		return [][]int{}
	}

	row, column := n, n
	top := 0
	bottom := row - 1
	left := 0
	right := column - 1

	clockArr := make([][]int, row)
	for i := range clockArr {
		clockArr[i] = make([]int, column)
	}

	num := 1

	for num <= row*column {
		// kanan
		for i := left; i <= right && num <= row*column; i++ {
			clockArr[top][i] = num
			num++
		}
		top++

		// bawah
		for i := top; i <= bottom && num <= row*column; i++ {
			clockArr[i][right] = num
			num++
		}
		right--

		// kiri
		for i := right; i >= left && num <= row*column; i-- {
			clockArr[bottom][i] = num
			num++
		}
		bottom--

		// atas
		for i := bottom; i >= top && num <= row*column; i-- {
			clockArr[i][left] = num
			num++
		}
		left++

	}
	return clockArr
}

func main() {
	fmt.Println(CreateSpiral(5))
}
