package main


func insertionSort(arr []int) []int {
	temp := 0
	for i := 1; i < len(arr); i++ {
		temp = arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}

	return arr
}


