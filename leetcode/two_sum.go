package main

func TwoSum(numbers []int, target int) [2]int {
	var res [2]int
	for i := 0; i < len(numbers); i++ {
		for j := (i + 1); j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return res
}
