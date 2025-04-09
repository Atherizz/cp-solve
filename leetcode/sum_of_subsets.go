package main

import "fmt"

func SumOfSubsets(set []int, target int) [][]int {
	var results [][]int
	var path []int
	sumRemaining := sum(set)
	recursiveSubset(0, 0, sumRemaining, set, target, &path, &results)

	return results
}

func recursiveSubset(index int, currSum int, sumRemaining int, set []int, target int, path *[]int, results *[][]int) {
	if currSum == target {
		solution := make([]int, len(*path))
		copy(solution, *path)
		*results = append(*results, solution)
		return
	}

	if index >= len(set) || currSum+sumRemaining < target || currSum > target {
		return
	}

	*path = append(*path, set[index])
	recursiveSubset(index+1, currSum+set[index], sumRemaining-set[index], set, target, path, results)
	*path = (*path)[:len(*path)-1]

	recursiveSubset(index+1, currSum, sumRemaining-set[index], set, target, path, results)
}

func sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
func main() {
	fmt.Println(SumOfSubsets([]int{5, 10, 12, 13, 15, 18}, 30))
}
