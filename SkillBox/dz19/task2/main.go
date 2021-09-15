package main

import "fmt"

const arrLength = 12

func find(arr [arrLength]int, number int) (index int) {
	index = -1
	for i := 0; i < arrLength; i++ {
		if arr[i] == number {
			index = i
			return
		}
	}

	return
}

func main() {
	fmt.Println("*** Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться) ***")

	var number int
	var arr = [arrLength]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print("Введите число для поиска: ")
	_, _ = fmt.Scan(&number)

	index := find(arr, number)

	if index >= 0 {
		fmt.Printf("Элемен найден на позиции: %d\n", index+1)
	} else {
		fmt.Println("Элемент в массиве не найден")
	}
}
