package main

import "fmt"

func removeDuplicates(nums []int) int {

	// if len(nums) == 0 {
    //     return 0
    // }

	j := 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
	}

	nums[j] = nums[len(nums) - 1]
	j++

	return j
}

// func removeDuplicates(nums []int) int {

// 	if len(nums) == 0 {
//         return 0
//     }

// 	j := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != nums[i-1] {
// 			nums[j] = nums[i]
// 			j++
// 		}
// 	}

// 	return j
// }

func main() {
	num := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(num)
}
