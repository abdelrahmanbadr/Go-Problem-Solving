package main

import "fmt"

func main() {
	
	fmt.Println(reverse(12345))
}

// number is 123
// number % 10 --> 3 && number /= 10 --> 12
// number % 10 --> 2 && number /= 10 --> 1
// number % 10 --> 1 && number /= 10 --> 0
func reverse(number int) int {

	reverse := 0
	unit := 1

	for number != 0 {
		reverse *= 10
		reverse += number % 10

		unit *= 10
		number /= 10

	}

	return reverse
}
