package main

import "fmt"

// [1,2,3,3,4,5,6]
// [3,3,6,7,9]
//[1,1,2,2,3,3,6,6,6]
//result should be [3,3,6]
func main() {
	arr1 := []int{1, 2, 3, 3, 4, 5, 6}
	arr2 := []int{3, 3, 6, 7, 9}
	arr3 := []int{1, 1, 2, 2, 3, 3, 6, 6, 6}
	findIntersectionTwo(arr1, arr2, arr3)
}
func findIntersectionTwo(arr1, arr2, arr3 []int) {
	var x, y, z int
	for x < len(arr1) && y < len(arr2) && z < len(arr3) {

		if arr1[x] == arr2[y] && arr1[x] == arr3[z] {
			fmt.Println(arr1[x])
			x++
			y++
			z++

		} else if arr1[x] < arr2[y] {
			x++
		} else if arr1[x] > arr2[y] {
			y++
		} else {
			z++
		}

	}

}
