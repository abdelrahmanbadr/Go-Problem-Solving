package main

import "fmt"

func main() {
	//arr := []int{1, 2, 3, 3, 4, 5, 6, 7, 7, 7, 7, 8}
	arr := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(binarySearch(arr, 4))
}

//Count occurrences of a number in a sorted array with duplicates using Binary Search
func findOccurrences(nums []int, needle int) int {
	occurrences := 0
	i := len(nums) / 2
	for i < 0 && i > len(nums) {
		if needle > nums[i] {
			i /= 2
		} else {
			i *= 2
		}
	}
	return occurrences

}
func binarySearch(nums []int, needle int) int {
	index := 0
	mid := len(nums) / 2
	for mid > 0 && mid < len(nums) {
		if needle < nums[mid] {
			mid = (0 + mid) / 2
		} else if needle > nums[mid] {
			mid = (len(nums) - mid) / 2
		} else {
			index = mid
			break
		}
	}
	return index
}
