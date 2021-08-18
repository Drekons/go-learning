package main

import (
	"fmt"
)

func main() {
	fmt.Println("*** Счастливый билет ***")

	var number int

	fmt.Print("Введите номер билета [1234]: ")
	fmt.Scan(&number)

	var firstPath int = number / 100
	var lastPath int = number % 100
	var firstPathElement1 int = firstPath / 10
	var firstPathElement2 int = firstPath % 10
	var lastPathElement1 int = lastPath / 10
	var lastPathElement2 int = lastPath % 10

	if firstPathElement1 == lastPathElement2 && firstPathElement2 == lastPathElement1 {
		fmt.Println("Билет зеркальный!!!")
	} else if firstPathElement1+firstPathElement2 == lastPathElement1+lastPathElement2 {
		fmt.Println("Билет счастливый!")
	} else {
		fmt.Println("Билет обычный :(")
	}
}
