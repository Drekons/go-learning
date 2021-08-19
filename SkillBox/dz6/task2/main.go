package main

import "fmt"

func main() {
	fmt.Println("*** Сумма двух чисел по единице ***")

	var numberOne int
	var numberTwo int

	fmt.Print("Введите первое число: ")
	_, _ = fmt.Scan(&numberOne)
	fmt.Print("Введите второе число: ")
	_, _ = fmt.Scan(&numberTwo)

	sum := numberOne + numberTwo

	for numberOne != sum {
		numberOne++
		fmt.Println(numberOne)
	}
}
