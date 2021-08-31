package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reversion(numbers []int) []int {
	var newNumbers []int

	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}

	return newNumbers
}

func main() {
	fmt.Println("*** Задание 2. Функция, реверсирующая массив ***")

	var numbers []int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rand.Intn(5)+5; i++ {
		numbers = append(numbers, rand.Intn(10))
	}

	fmt.Printf("Исходный массив: %v\n", numbers)
	fmt.Printf("Реверсивный массив: %v\n", reversion(numbers))
}
