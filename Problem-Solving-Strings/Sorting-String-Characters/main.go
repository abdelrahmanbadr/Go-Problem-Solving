package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(sortCharsOfString("zkjhgAfdsa"))
}

func sortCharsOfString(str string) string {

	str = strings.ToLower(str)
	letters := [256]int{}
	for _, letter := range str {
		letters[letter]++
	}
	sortedString := ""
	for letter, occurance := range letters {

		for i := 0; i < occurance; i++ {
			sortedString = sortedString + string(letter)
		}
	}
	return sortedString

}
