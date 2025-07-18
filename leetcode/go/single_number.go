package main

// import "fmt"

func singleNumber(nums []int) int {
	data := make(map[int]int)

	for _, num := range nums {
		data[num]++
	}

	for _, num := range nums {
		if data[num] == 1 {
			return num
		}
	}

	return 0
}

// func main() {
// 	fmt.Println(singleNumber([]int{4,1,2,1,2}))
// }
