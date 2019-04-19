package main

import "fmt"

//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//Input: nums = [4,5,6,7,0,1,2], target = 0 , Output: 4
//Input: nums = [4,5,6,7,0,1,2], target = 3 , Output: -1
func main() {
	nums := []int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	fmt.Println(search(nums, 5))
}
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	point := findPivotIndex(nums)

	if nums[point] == target {
		return point
	}
	var min, max int
	if target > nums[0] {
		min = 0
		max = point - 1
	} else if target < nums[0] {
		min = point + 1
		max = len(nums) - 1
	} else {
		return 0
	}

	for max >= min {
		mid := (min + max) / 2

		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func findPivotIndex(nums []int) int {
	min, max := 0, len(nums)-1

	for max > min {

		mid := (min + max) / 2

		if nums[mid] > nums[min] {
			min = mid
		} else {
			max = mid
		}

	}
	return max
}
