package main

import "fmt"

func GetIndicesOfItemWeights(arr []int, limit int) (int, int) {
	// var wightsMap map[int]int //will not work because it's = nil now
	wightsMap := map[int]int{}
	for key, value := range arr {
		complement := limit - value

		if complementIndex, ok := wightsMap[complement]; ok {
			return complementIndex, key
		}
		wightsMap[value] = key
	}
	return 0, 0
}

func main() {
	fmt.Println(GetIndicesOfItemWeights([]int{4, 6, 10, 15, 16}, 21))
}
