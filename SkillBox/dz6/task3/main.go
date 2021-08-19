package main

import "fmt"

func main() {
	fmt.Println("*** Расчёт суммы скидки ***")

	var price float64
	var discountPresent float64
	maxDiscountCost := 2000.
	maxDiscountPresent := 30.

	fmt.Print("Введите цену товара: ")
	_, _ = fmt.Scan(&price)
	fmt.Print("Введите скидку (%): ")
	_, _ = fmt.Scan(&discountPresent)

	sale := price * discountPresent / 100
	for sale > maxDiscountCost || discountPresent > maxDiscountPresent {
		discountPresent--
		sale = price * discountPresent / 100
	}

	fmt.Println("Сумма скидки:", sale)
}
