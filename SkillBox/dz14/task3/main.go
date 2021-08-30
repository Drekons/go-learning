package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multiply(x int) (res int) {
	res = x * rand.Intn(10)
	return
}
func plus(x int) (res int) {
	res = x + rand.Intn(10)
	return
}

func modifyNumber(x int) (res int) {
	res = multiply(x)
	res = plus(res)
	return
}

func main() {
	fmt.Println("*** Задание 3. Именованные возвращаемые значения ***")

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Итоговое число: ", modifyNumber(rand.Intn(10)))
}
