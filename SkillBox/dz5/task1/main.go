package main

import "fmt"

func main() {
	fmt.Println("*** Определение координатной плоскости точки ***")

	var x int
	var y int
	var position int = 0

	fmt.Print("Введите координату x: ")
	fmt.Scan(&x)

	fmt.Print("Введите координату y: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		position = 1
	} else if x < 0 && y > 0 {
		position = 2
	} else if x < 0 && y > 0 {
		position = 2
	} else if x < 0 && y < 0 {
		position = 3
	} else if x > 0 && y < 0 {
		position = 4
	}

	fmt.Println("Позиция в координатной четверти", position)
}
