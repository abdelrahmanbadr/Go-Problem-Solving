package main

import (
	"fmt"
	"math"
)

func Root(x float64, n float64) float64 {
	var start float64
	end := x
	mid := (start + end) / 2
	//for start <= end {
	for (mid - start) >= 0.0001 { // for approximate values for 0.0XX numbers

		if math.Pow(mid, n) > x {
			end = mid
		} else if math.Pow(mid, n) < x {
			start = mid
		} else {
			break
		}
		mid = (start + end) / 2
	}
	return mid
}

func main() {
	fmt.Println(Root(64, 2))
	fmt.Println(Root(81, 4))
	fmt.Println(Root(7, 3))
}
