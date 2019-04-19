package main

import "fmt"

// AmexPrblem [20,10,5,1] the number is 53

func getMinimumNumberOfCoins(arr []int, number int) []int {
	result := []int{}
	for _, value := range arr {

		for number >= value {
			number -= value
			result = append(result, value)
		}
		if number == 0 {
			break
		}
	}
	return result
}

func main() {
	x := "xooo"
	x = ^x
	fmt.Println(x)
	result := []int{20, 10, 5, 1}
	fmt.Println(getMinimumNumberOfCoins(result, 2))
}
