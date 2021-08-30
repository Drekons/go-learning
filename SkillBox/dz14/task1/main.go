package main

import "fmt"

func isEven(number int) bool {
	return number%2 == 0
}

func main() {
	fmt.Println("*** Задание 1. Функция возвращающая результат ***")

	var number int

	fmt.Print("Введите число: ")
	_, _ = fmt.Scan(&number)

	if isEven(number) {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число не чётное")
	}

}
