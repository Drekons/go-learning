package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("*** Угадай число ***")

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	number := rand.Intn(101)

	fmt.Println("Программа загадала число, попробуйте его угадать", number)

	for {
		var guess int
		for guess < min || guess > max {
			fmt.Print("Введите число от ", min, " до ", max, ": ")
			fmt.Scan(&guess)
		}

		if guess == number {
			fmt.Println("Вы угадали! Это число", number)
			break
		}

		fmt.Println("Не верно!")

		if guess > number {
			fmt.Println("Загаданное число меньше")
		} else {
			fmt.Println("Загаданное число больше")
		}
	}

}
