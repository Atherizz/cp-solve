package main



func FindClosestNumber(nums []int) int {
	gap := make([]int, len(nums))
	for i, num := range nums {
		if num < 0 {
			gap[i] = num * -1
		} else {
			gap[i] = num
		}
	}
	closest := gap[0]
	index := 0
	for i, num := range gap {
		if num < closest || num == closest && nums[i] > nums[index] {
			closest = num
			index = i
		}
	}
	return nums[index]
}

