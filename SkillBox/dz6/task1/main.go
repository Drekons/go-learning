package main

import "fmt"

func main() {
	fmt.Println("*** Написание простого цикла ***")

	var number int

	fmt.Print("Введите произвольное число: ")
	_, _ = fmt.Scan(&number)

	for i := 0; i < number; i++ {
		fmt.Println(i)
	}
}
