package main

import "fmt"

func main() {
	fmt.Println(twoMissingNumbers_better([]int{1, 5, 3, 6}))

}

//Given array contain all the numbers between 1 and n except two numbers
// [2,4,3] => 1, 5
//O(n) time complexity and O(n) Extra Space
func twoMissingNumbers(arr []int) [2]int {
	min := 1
	max := len(arr) + 2

	mp := make(map[int]bool, max)
	for _, val := range arr {
		mp[val] = true
	}
	missing := [2]int{}
	j := 0
	for i := min; i <= max; i++ {
		if _, ok := mp[i]; !ok {
			missing[j] = i
			j++
			if j == 2 {
				break
			}
		}
	}
	return missing
}

//O(n) time complexity and O(1) Extra Space
//https://www.geeksforgeeks.org/find-two-missing-numbers-set-1-an-interesting-linear-time-solution/
//https://www.geeksforgeeks.org/find-two-missing-numbers-set-2-xor-based-solution/

/**
Input : 1 3 5 6, n = 6
Sum of missing integers = n*(n+1)/2 - (1+3+5+6)
						= 21 - 15 = 6

Average of missing integers = 6/2 = 3.

Sum of array elements less than or equal to average = 1 + 3 = 4
Sum of array elements greater than average = 5 + 6 = 11

Sum of natural numbers from 1 to avg = avg*(avg + 1)/2
                                     = 3*4/2 = 6
Sum of natural numbers from avg+1 to n
									= Sum of array elements - Sum of natural numbers from 1 to avg
									=   n*(n+1)/2 - avg*(avg + 1)/2
									=  21 - 6
									=  15

First missing number = Sum of natural numbers 1 to avg -  Sum of array elements less than or equal to avg
                     = 6 - 4 = 2

Second missing number = Sum of natural numbers from avg+1 to n - Sum of array elements greater than avg
					  = 15 - 11 = 4
**/

func twoMissingNumbers_better(arr []int) [2]int {

	arraySize := len(arr) + 2
	arraySum := arraySize * (arraySize + 1) / 2

	var avg, sumSmallerHalf, sumGreaterHalf int
	avg = (arraySum - getSum(arr)) / 2

	for _, val := range arr {
		if val <= avg {
			sumSmallerHalf += val
			continue
		}
		sumGreaterHalf += val
	}
	sumFromStartToAvg := avg * (avg + 1) / 2
	sumFromAvgtToEnd := arraySum - sumFromStartToAvg
	fmt.Println(arraySum, sumFromStartToAvg)
	return [2]int{sumFromStartToAvg - sumSmallerHalf, sumFromAvgtToEnd - sumGreaterHalf}

}

func getSum(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}
