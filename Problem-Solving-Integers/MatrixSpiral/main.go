package main

import "fmt"

var result []int

func SpiralCopy(inputMatrix [][]int) []int {
	if len(inputMatrix) == 0 {
		return result
	}

	rowsNumber := len(inputMatrix)
	colsNumber := len(inputMatrix[0])
	var leftCol, topRow int
	rightCol := colsNumber - 1
	bottomRow := rowsNumber - 1
	for leftCol <= rightCol && topRow <= bottomRow {

		for i := leftCol; i <= rightCol; i++ {
			addToResult(inputMatrix[topRow][i])
		}
		topRow++

		for i := topRow; i <= bottomRow; i++ {
			addToResult(inputMatrix[i][rightCol])
		}
		rightCol--

		for i := rightCol; i >= leftCol; i-- {
			addToResult(inputMatrix[bottomRow][i])
		}
		bottomRow--

		for i := bottomRow; i >= topRow; i-- {
			addToResult(inputMatrix[i][leftCol])
		}
		leftCol++

	}

	return result
}

func addToResult(number int) {
	result = append(result, number)
}
func main() {
	inputMatrix := [][]int{
		{1, 2, 3, 4, 5},
		{14, 15, 16, 17, 6},
		{13, 20, 19, 18, 7},
		{12, 11, 10, 9, 8}}

	SpiralCopy(inputMatrix)
	fmt.Println(result)
}
