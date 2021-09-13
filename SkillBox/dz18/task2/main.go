package main

import (
	"fmt"
	"math/rand"
	"time"
)

func applyFunc(x, y int, A func(int, int) int) {
	defer func() {
		fmt.Printf("Результат: %d\n", A(x, y))
	}()

	fmt.Printf("Значения x: %d y: %d\n", x, y)
}

func main() {
	fmt.Println("*** Задание 2. Анонимные функции  ***")
	rand.Seed(time.Now().UnixNano())

	fmt.Println("-----------")

	applyFunc(rand.Intn(100), rand.Intn(100), func(x int, y int) int {
		fmt.Println("Операция умножение")
		return x * y
	})

	fmt.Println("-----------")

	applyFunc(rand.Intn(100), rand.Intn(100), func(x int, y int) int {
		fmt.Println("Операция вычитание")
		return x - y
	})

	fmt.Println("-----------")

	applyFunc(rand.Intn(100), rand.Intn(100), func(x int, y int) int {
		fmt.Println("Операция сложение")
		return x + y
	})
}
