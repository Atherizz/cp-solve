package main

func GreedyKnapsack(capacity int, items [][]int) []float64 {
	avgProfit := make([][2]int, len(items))

	for i := 0; i < len(items); i++ {
		avgProfit[i] = [2]int{items[i][1] / items[i][0], i}
	}

	insertionSort(avgProfit, items)

	quantities := make([]float64, len(items))

	for i := 0; i < len(items) && capacity > 0; i++ {
		idx := avgProfit[i][1] 
		if capacity >= items[i][0] {
			quantities[idx] = 1
			capacity -= items[i][0]
		} else {
			fraction := float64(capacity) / float64(items[i][0])
			quantities[idx] = fraction
			capacity = 0
			break
		}
	}


	return quantities

}

func insertionSort(profitIndex [][2]int, items [][]int) {
	for i := 1; i < len(profitIndex); i++ {
		temp := profitIndex[i]
		itemTemp := items[i]

		j := i - 1

		for j >= 0 && profitIndex[j][0] < temp[0] {
			profitIndex[j+1] = profitIndex[j]
			items[j+1] = items[j]
			j--
		}
		profitIndex[j+1] = temp
		items[j+1] = itemTemp
	}

}

