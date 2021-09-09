package main

import "fmt"

const FirstRows = 3
const FirstCols = 5
const SecondRows = 5
const SecondCols = 4

func multiply(
	firstMatrix [FirstRows][FirstCols]int,
	secondMatrix [SecondRows][SecondCols]int,
) [FirstRows][SecondCols]int {
	var resultMatrix [FirstRows][SecondCols]int

	for i := 0; i < FirstRows; i++ {
		for j := 0; j < SecondCols; j++ {
			for k := 0; k < FirstCols; k++ {
				resultMatrix[i][j] += firstMatrix[i][k] * secondMatrix[k][j]
			}
		}
	}

	return resultMatrix
}

func main() {
	fmt.Println("Задание 2. Умножение матриц")

	firstMatrix := [FirstRows][FirstCols]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}
	secondMatrix := [SecondRows][SecondCols]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	fmt.Printf("Умножение матрицы %v\n", firstMatrix)
	fmt.Printf("на матрицу %v\n", secondMatrix)
	fmt.Printf("даст матрицу %v\n", multiply(firstMatrix, secondMatrix))

}
