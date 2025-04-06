package main

func removeElement(nums []int, val int) int {
	j := 0
	for i, _ := range nums {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

