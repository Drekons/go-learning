package main

import "fmt"

const arrLength = 12

func find(arr [arrLength]int, number int) (index int) {
	index = -1
	min := 0
	max := arrLength - 1
	for max >= min {
		middle := (max + min) / 2
		if arr[middle] == number {
			for arr[middle] == number {
				index = middle
				middle--
			}
			return
		} else if arr[middle] > number {
			max = middle - 1
		} else {
			min = middle + 1
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
