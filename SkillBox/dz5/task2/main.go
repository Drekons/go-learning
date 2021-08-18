package main

import "fmt"

func main() {
	fmt.Println("*** Проверить, что одно из чисел положительное ***")

	var a int
	var b int
	var c int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Print("Введите третье число: ")
	fmt.Scan(&c)

	if a > 0 || b > 0 || c > 0 {
		fmt.Println("Есть положительное число")
	} else {
		fmt.Println("Все числа отрицательные")
	}
}
