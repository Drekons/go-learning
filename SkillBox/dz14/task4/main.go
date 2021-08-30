package main

import "fmt"

const ConstantTen = 10
const ConstantHundred = 100

func modifyTen(x int) int {
	return x + ConstantTen
}

func modifyOneHundredTen(x int) int {
	return x + ConstantTen + ConstantHundred
}

func modifyOneHundred(x int) int {
	return x + ConstantHundred
}

func main() {
	fmt.Println("*** Задание 4. Область видимости переменных ***")

	var number int

	fmt.Print("Введите число: ")
	_, _ = fmt.Scan(&number)

	number = modifyTen(number)
	number = modifyOneHundredTen(number)
	number = modifyOneHundred(number)

	fmt.Println("После преобразования число равно", number)
}
