package main

import "fmt"

func main() {
	fmt.Println("*** Шахматная доска ***")

	var horizontalCount int
	var verticalCount int

	fmt.Print("Введите количество клеток по горизонтали: ")
	_, _ = fmt.Scan(&horizontalCount)

	fmt.Print("Введите количество клеток по вертикали: ")
	_, _ = fmt.Scan(&verticalCount)

	for i := 0; i < verticalCount; i++ {
		for j := 0; j < horizontalCount; j++ {
			if j%2 == i%2 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}
