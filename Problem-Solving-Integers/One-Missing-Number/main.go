package main

import "fmt"

func main() {
	xorMissingNumber([]int{4, 3, 2, 5})
	//fmt.Println()
	//fmt.Println(missingNumber([]int{1, 3, 4, 5}))
}

func missingNumber(arr []int) int {
	naturalArraySize := len(arr) + 1
	naturalNumbersSum := naturalArraySize * (naturalArraySize + 1) / 2
	arrSum := 0
	for _, val := range arr {
		arrSum += val
	}
	return naturalNumbersSum - arrSum
}

// when u xor number with itself it results zero
// so if input is : [1,3]
// 1 ^ 3 ^ 1 ^ 2 ^ 3 = 2
func xorMissingNumber(arr []int) int {

	xor := 0
	for _, val := range arr {
		xor ^= val
	}
	for i := 1; i <= len(arr)+1; i++ {
		xor ^= i
	}

	return xor
}

// when u xor number with itself it results zero -->ex: 3 xor 3 = 0
func xor(num1, num2 int) {
	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	fmt.Println(num1, num2)
	//return (num1 | num2) && !(num1 && num2)
}

func swap() {
	//$x = 1;
	//$y = 2;
	//
	//$x = $y + $x;
	//$y = $x - $y;
	//$x = $x - $y;
	//
	//$str1 = "x";
	//$str2 = "y";
	//$str1 = ($str1 ^ $str2);
	//$str2 = ($str2 ^ $str1);
	//$str1 = ($str1 ^ $str2);
}

//xor
