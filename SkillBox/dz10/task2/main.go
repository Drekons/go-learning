package main

import (
	"fmt"
	"math"
)

func main() {
	var account int
	var capitalizationProc int
	var years int

	fmt.Print("Введите сумму которую хотите положить на счёт: ")
	_, _ = fmt.Scan(&account)
	fmt.Print("Введите процент капитализации: ")
	_, _ = fmt.Scan(&capitalizationProc)
	fmt.Print("Введите на какое количество лет будет вклад: ")
	_, _ = fmt.Scan(&years)

	var bank = 0.
	var result = float64(account)
	var capitalization = float64(capitalizationProc) / 100.

	for i := 1; i <= years*12; i++ {
		add := result * capitalization
		addClean := math.Floor(add*100) / 100
		result += addClean

		bank += add - addClean
	}

	fmt.Printf("Вы получите: %f\nБанк получит: %f\n", result, bank)
}
