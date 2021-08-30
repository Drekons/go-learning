package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomCors() (int, int, int, int, int, int) {
	return rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)
}

func cordModify(x int, y int) (int, int) {
	return 2*x + 10, -3*y - 5
}

func main() {
	fmt.Println("*** Задание 2. Функция возвращающая несколько значений ***")

	rand.Seed(time.Now().UnixNano())

	x1, y1, x2, y2, x3, y3 := getRandomCors()

	x, y := cordModify(x1, y1)
	fmt.Printf("x1: %d, y1: %d | modify x: %d, y: %d\n", x1, y1, x, y)
	x, y = cordModify(x2, y2)
	fmt.Printf("x2: %d, y2: %d | modify x: %d, y: %d\n", x1, y1, x, y)
	x, y = cordModify(x3, y3)
	fmt.Printf("x3: %d, y3: %d | modify x: %d, y: %d\n", x1, y1, x, y)
}
