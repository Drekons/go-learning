package main

import "fmt"

const ROWS = 3
const COLS = 3

func det(matrix [ROWS][COLS]int) int {
	var det int

	for r0 := 0; r0 < ROWS; r0++ {
		r1, r2 := getRightCords(r0)
		det += matrix[0][r0] * matrix[1][r1] * matrix[2][r2]
	}

	for r0 := ROWS - 1; r0 >= 0; r0-- {
		r1, r2 := getLeftCords(r0)
		det -= matrix[0][r0] * matrix[1][r1] * matrix[2][r2]
	}

	return det
}

func getRightCords(r0 int) (r1 int, r2 int) {
	r1 = r0 + 1
	r2 = r1 + 1

	if r1 >= ROWS {
		r1 = r1 - ROWS
	}

	if r2 >= ROWS {
		r2 = r2 - ROWS
	}

	return
}

func getLeftCords(r0 int) (r1 int, r2 int) {
	r1 = r0 - 1
	r2 = r1 - 1

	if r1 < 0 {
		r1 = r1 + ROWS
	}

	if r2 < 0 {
		r2 = r2 + ROWS
	}

	return
}

func main() {
	fmt.Println("Задание 1. Подсчёт определителя")

	matrix := [ROWS][COLS]int{
		{1, 3, 3},
		{5, 5, 5},
		{2, 2, 1},
	}

	fmt.Printf("Определитель для матрицы %v : %d\n", matrix, det(matrix))

}
