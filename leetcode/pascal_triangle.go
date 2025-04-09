package main

func generate(numRows int) [][]int {
	pascal := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		pascal[i] = make([]int, i+1)
		pascal[i][0], pascal[i][i] = 1, 1
		for j := 1; j < i; j++ {
			pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]
		}
	}
	return pascal
}

