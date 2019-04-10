package main

import "fmt"

//test cases :
// abcabcbb --> 3
// abba --> 2
// pwwkew -> 3
func main() {
	//fmt.Println(lengthOfLongestSubstring("xoxabadoxoxo"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) //should return 2
}

func lengthOfLongestSubstring(input string) int {
	containsChar := make(map[int32]int)
	maxLength, start := 0, 0
	for i, val := range input {
		if _, ok := containsChar[val]; ok {
			// to make sure start is not going to old index
			if start < containsChar[val]+1 {
				start = containsChar[val] + 1
			}
		}
		containsChar[val] = i
		maxLength = Max(maxLength, i+1-start)
	}
	return maxLength
}

func Max(number1 int, number2 int) int {

	if number1 > number2 {
		return number1
	}
	return number2
}
