package main

import "fmt"

//For example, if the array is [3, 5, 2, -4, 8, 11] and the sum is 7,
//your program should return [[11, -4], [2, 5]] because 11 + -4 = 7 and 2 + 5 = 7.
func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 7, 4}, 6))
}

func twoSum(arr []int, target int) []int {
	mp := make(map[int]bool, len(arr))
	for _, v := range arr {
		complement := target - v
		if _, ok := mp[complement]; ok {
			return []int{complement, v}
		}
		mp[v] = true
	}
	return []int{}
}
