package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] == target {
			return mid
		}
	}

	if nums[mid] > target {
		return mid
	}
	return mid + 1
}
