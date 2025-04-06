package main

import "fmt"

func knapsack(n, capacity int, weights, values []int) [][]int {
	table := make([][]int, n+1)
	maks := 0

	for i := 0; i < n+1; i++ {
		table[i] = make([]int, (capacity + 1))
	}

	for i := 1; i < len(table); i++ {
		for j := 1; j < len(table[i]); j++ {
			if i == 1 && j >= weights[i-1] {
				table[i][j] = values[i-1]
			} else {
				if j < weights[i-1] {
					table[i][j] = table[i-1][j]
				} else {
					a := values[i-1] + table[i-1][j-weights[i-1]]
					b := table[i-1][j]
					if a > b {
						maks = a
					} else {
						maks = b
					}
					table[i][j] = maks
				}
			}
		}
	}

	return table
}

func getMaxProfit(weights, values []int, table [][]int) [][]int {
	n := len(table) - 1
	capacity := len(table[0]) - 1
	maxProfit := 0
	capRemain := capacity
	i := n
	j := capacity

	res := make([]int, n)

	for i > 0 && j > 0 {
		if table[i][j] != table[i-1][j] {
			maxProfit += values[i-1]
			res[i-1] = 1
			capRemain -= weights[i-1]
			j -= weights[i-1]
		}
		i--
	}

	if i == 0 && j >= weights[0] && table[0][j] > 0 {
		maxProfit += values[0]
	}

	return [][]int{[]int{maxProfit}, res}
}

func main() {
	table := knapsack(4, 8, []int{3, 2, 5, 4}, []int{1, 2, 5, 6})
	fmt.Println(table)
	fmt.Println(getMaxProfit([]int{3, 2, 5, 4}, []int{1, 2, 5, 6}, table))

}
