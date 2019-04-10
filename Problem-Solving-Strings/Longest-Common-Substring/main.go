package main

import "fmt"

func main() {
	fmt.Println(LongestCommonSubStringLength("GeeksforGeeks", "GeeksQuiz"))
}

// Input : X = "GeeksforGeeks", y = "GeeksQuiz"  Output : 5
func LongestCommonSubStringLength(str1, str2 string) int {

	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	var max int
	cache := make([][]int, len(str1))

	for i, str1Char := range str1 {
		cache[i] = make([]int, len(str2))
		for j, str2Char := range str2 {
			if str1Char == str2Char {
				if i == 0 || j == 0 {
					cache[i][j] = 1
				} else {
					cache[i][j] = cache[i-1][j-1] + 1
				}
			}
			max = Max(max, cache[i][j])

		}
	}

	return max
}

func Max(number1 int, number2 int) int {

	if number1 > number2 {
		return number1
	}
	return number2
}
