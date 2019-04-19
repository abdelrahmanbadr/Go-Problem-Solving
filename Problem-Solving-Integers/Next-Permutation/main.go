package main

import (
	"fmt"
)

// level : Hard
func main() {

	nextPermutation([]int{1, 5, 8, 4, 7, 6, 5, 3, 1}) //15 01269
	//nextPermutation([]int{1, 5, 8, 4, 7, 6, 5, 3, 1}) //1585 13467
	//nextPermutation([]int{1, 2, 9, 6, 5})             //15269
}

// 158476531 --> 1585 13467
// 12965      --> 15 269
//	[1, 2, 9, 6, 5, 1, 0] --> [1,5,0,1,2,6,9]
func nextPermutation(nums []int) {
	fmt.Println(nums)
	i := len(nums) - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}

	x := findNextGreaterNumber(nums[i:len(nums)], nums[i-1])

	if i > 0 {
		swap(&nums[i-1], &nums[x+i])
	}
	BubbleSort(nums[i:len(nums)])

	fmt.Println(nums)

}
func findNextGreaterNumber(nums []int, comp int) int {

	smallestIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[smallestIndex] > nums[i] && nums[i] > comp {
			smallestIndex = i
		}
	}
	return smallestIndex
}

func BubbleSort(items []int) {

	L := len(items)

	for i := 0; i < L; i++ {

		for j := 0; j < (L - 1 - i); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}

}

func swap(x, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}
