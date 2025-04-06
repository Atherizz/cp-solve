package main

import "fmt"

func getRow(rowIndex int) []int {

	pascal := make([][]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		pascal[i] = make([]int, i+1)
		pascal[i][0], pascal[i][i] = 1, 1
		for j := 1; j < i; j++ {
			pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]

		}
	}
	return pascal[rowIndex]

}

func main() {
	fmt.Println(getRow(4))
}
