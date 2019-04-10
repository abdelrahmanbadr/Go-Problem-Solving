package main

import "fmt"

type Stack struct {
	StackArray []int
	Top        int
}

func (s *Stack) IsEmpty() bool {
	return s.Top == 0
}

func (s *Stack) Push(v int) {
	s.StackArray = append(s.StackArray, v)
	s.Top++
}

func (s *Stack) Pop() int {

	lastValue := s.StackArray[s.Top-1]
	//remove last element
	s.StackArray = s.StackArray[:s.Top-1]
	s.Top--
	return lastValue
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(2)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	// Output:
	// 2
	// 1
	// true

}
