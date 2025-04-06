package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		middle := (low + high) / 2
		if nums[middle] == target {
			return middle
		}
		if target > nums[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if i < len(nums) - 1 && nums[i] < target && nums[i+1] > target {
			return i+1
		} 
	}

	if nums[len(nums) - 1] < target {
		return len(nums)
	}

	if nums[0] > target {
		return 0
	}

	return -1
}
