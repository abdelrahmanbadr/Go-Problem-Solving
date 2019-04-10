package main

import "fmt"

func main() {
	stringPermutations("ab")

}

//see swap diagram here : https://leetcode.com/articles/short-permutation-in-a-long-string/
func stringPermutations(str string) {
	result := []string{}
	permutations(&result, "", str)
	fmt.Println(result)
}

func permutations(result *[]string, prefix, suffix string) {

	if suffix == "" {
		*result = append(*result, prefix)
	} else {
		for i, val := range suffix {
			newPrefix := prefix + string(val)
			newSuffix := string(suffix[0:i]) + string(suffix[i+1:len(suffix)])
			permutations(result, newPrefix, newSuffix)
		}
	}

}
