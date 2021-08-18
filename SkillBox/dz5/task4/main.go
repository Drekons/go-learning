package main

import "fmt"

func main() {
	fmt.Println("*** Сумма без сдачи ***")

	var price int
	var a int
	var b int
	var c int
	var canPay = false

	fmt.Print("Введите стоимость товара: ")
	fmt.Scan(&price)

	fmt.Print("Введите номинал первой монеты: ")
	fmt.Scan(&a)

	fmt.Print("Введите номинал второй монеты: ")
	fmt.Scan(&b)

	fmt.Print("Введите номинал третей монеты: ")
	fmt.Scan(&c)

	if a+b+c == price {
		canPay = true
	} else if a+b == price || a+c == price || b+c == price {
		canPay = true
	}

	if canPay {
		fmt.Println("Можно оплатить без сдачи")
	} else {
		fmt.Println("Нельзя оплатить без сдачи")
	}
}
