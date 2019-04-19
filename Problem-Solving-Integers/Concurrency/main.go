package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 2

	close(c)
	_, y := <-c
	val, x := <-c
	fmt.Println(y)      // true  even if the channel is closed u can receive data from it
	fmt.Println(val, x) // 0 ,false  because not valid data (no thing in the channel)

}
