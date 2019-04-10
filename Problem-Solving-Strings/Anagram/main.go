package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAnagram("nabana", "banAna"))
}

//anagram words like (banana and nabana) or (baba bbaa)
func isAnagram(str1 string, str2 string) bool {

	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	if len(str1) != len(str2) {
		return false
	}
	letters := make(map[int32]int)
	for _, letter := range str1 {
		letters[letter]++
	}
	for _, letter := range str2 {
		letters[letter]--
	}
	for _, letter := range letters {
		if letter != 0 {
			return false
		}
	}
	return true
}
