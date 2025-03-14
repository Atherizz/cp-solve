package main


func FindEvenIndex(arr []int) int {
	sumLeft := make([]int, len(arr))
	sumRight := make([]int, len(arr))
	for i, _ := range arr {
			for j := (i - 1); j >= 0; j-- {
				sumLeft[i] += arr[j]
			}
			for j := (i + 1); j < len(arr); j++ {
				sumRight[i] += arr[j]
			}
			if sumLeft[i] == sumRight[i] {
				return i
			}
	}

	return -1

}
