package main


type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	res := PosPeaks{}
	for i := 1; i < (len(array) - 1); i++ {
		if array[i] > array[i-1] {
			j := i
			for j+1 < len(array) && array[j] == array[j+1] {
				j++
			}

			if j+1 < len(array) && array[j] > array[j+1] {
				res.Pos = append(res.Pos, i)
				res.Peaks = append(res.Peaks, array[i])
			}
		}
	}
	return res
}
