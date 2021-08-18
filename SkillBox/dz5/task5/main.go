package main

import "fmt"

func main() {
	fmt.Println("*** Определение максимальных процентов ***")

	var rateA int
	var rateB int
	var rateC int

	fmt.Print("Введите первую ставку: ")
	fmt.Scan(&rateA)

	fmt.Print("Введите вторую ставку: ")
	fmt.Scan(&rateB)

	fmt.Print("Введите третью ставку: ")
	fmt.Scan(&rateC)

	if rateA >= rateC && rateA >= rateB {
		fmt.Println("Выгодные ставки:", rateB, "и", rateC)
	} else if rateB >= rateC && rateB >= rateA {
		fmt.Println("Выгодные ставки:", rateA, "и", rateC)
	} else {
		fmt.Println("Выгодные ставки:", rateA, "и", rateB)
	}
}
