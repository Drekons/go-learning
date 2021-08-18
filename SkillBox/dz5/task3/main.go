package main

import "fmt"

func main() {
	fmt.Println("*** Проверить, что есть совпадающие числа ***")

	var a int
	var b int
	var c int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	fmt.Print("Введите третье число: ")
	fmt.Scan(&c)

	if a == b || b == c || a == c {
		fmt.Println("Есть совпадающие числа")
	} else {
		fmt.Println("Все числа уникальные")
	}
}
