package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SumDifRev(n int) int {
	data := make(map[int]int)
	counter := 0

	for i := 10; i < int(^uint(0)>>1); i++ {
		res := []string{}
		str := strconv.Itoa(i)
		for j := len(str) - 1; j >= 0; j-- {
			res = append(res, string(str[j]))
		}
		resStr := strings.Join(res, "")
		num, err := strconv.Atoi(resStr)
		if err != nil {
			continue
		}
		sum := i + num
		substract := i - num

		if substract != 0 && sum%AbsInt(substract) == 0 {
			counter++
			data[counter] = i
			if counter == n {
				break
			}
		}
	}

	return data[counter]

}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Print(SumDifRev(3))
}
