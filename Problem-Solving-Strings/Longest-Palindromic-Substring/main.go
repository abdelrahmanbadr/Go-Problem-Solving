package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := ""
	fmt.Scanf("%s", &input)
	//fmt.Println(input)
	fmt.Println(longestPalindrome(input))
	//fmt.Println(isPalindrome("abbba"))
}

//@todo implement again
func longestPalindrome(input string) string {

	answer := ""
	inputLength := len(input)
	for i := 0; i < inputLength; i++ {
		for j := i; j < inputLength+1; j++ {
			if i < j {
				subString := input[i:j]
				if isPalindrome(subString) && len(answer) < j-i {
					answer = subString
				}
			}
		}
	}
	return answer
}
func longestPalindrome_betterRunTime(s string) string {
	if len(s) == 1 {
		return s
	}
	ans := ""
	ansLength := 0

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i+ansLength; j-- {
			subString := s[i:j]
			if isPalindrome(subString) && len(subString) > ansLength {
				ans = subString
				ansLength = len(subString)
			}
		}
	}

	return ans
}
func isPalindrome(input string) bool {

	inputLength := len(input)
	str := strconv.Itoa(123456)
	for i := 0; i < inputLength/2; i++ {
		if input[i] != input[inputLength-i-1] {
			return false
		}
	}

	return true
}
