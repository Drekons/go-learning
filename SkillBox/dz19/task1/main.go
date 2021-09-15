package main

import (
	"fmt"
	"math/rand"
	"time"
)

const arrLength = 10

func findCountOddEvenAfterNumberInArray(number int, arr [arrLength]int) (index, even, odd int) {
	index = -1
	for i := 0; i < arrLength; i++ {
		if index >= 0 {
			if arr[i]%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		if arr[i] == number && index < 0 {
			index = i
		}
	}
	return
}

func main() {
	fmt.Println("*** Задание 1. Подсчёт чётных и нечётных чисел в массиве ***")

	rand.Seed(time.Now().UnixNano())

	var arr [arrLength]int
	var number int

	for i := 0; i < arrLength; i++ {
		arr[i] = rand.Intn(10 + i)
	}

	fmt.Printf("Исходный массив %v\n", arr)

	fmt.Print("Введите число для поиска: ")
	_, _ = fmt.Scan(&number)

	index, even, odd := findCountOddEvenAfterNumberInArray(number, arr)

	fmt.Printf(
		"Позиция %d, кол-во чётных после позиции %d, кол-во не чётных после позиции %d\n",
		index+1, even, odd,
	)
}
