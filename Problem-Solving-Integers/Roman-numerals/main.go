package main

import "fmt"

func main() {

	fmt.Println(romanToInt("IV"))
}
func romanToInt(s string) int {

	total := 0
	for i, char := range s {
		if i != len(s)-1 {
			total += romanInt(string(char), string(s[i+1]))
		} else {
			total += romanInt(string(char), "")
		}

	}
	return total

}
func romanInt(char string, next string) int {

	mp := make(map[string]int)
	mp["I"] = 1
	mp["V"] = 5
	mp["X"] = 10
	mp["L"] = 50
	mp["C"] = 100
	mp["D"] = 500
	mp["M"] = 1000
	if next == "" {
		return mp[char]
	}
	fmt.Println()
	if mp[char]*5 == mp[next] || mp[char]*10 == mp[next] {

		return mp[char] * -1
	}
	return mp[char]
}
