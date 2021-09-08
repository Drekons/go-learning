package main

import "fmt"

const arrayLen = 6

func bubbleSort(array [arrayLen]int) [arrayLen]int {
	for i := 0; i < arrayLen-1; i++ {
		for j := 0; j < arrayLen-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

func main() {
	fmt.Println("Задание 2. Сортировка пузырьком")

	var arrayForSort [arrayLen]int

	for i := 0; i < arrayLen; i++ {
		fmt.Printf("Введите элемент массива %d: ", i+1)
		_, _ = fmt.Scan(&arrayForSort[i])
	}

	fmt.Printf("Полученный массив: %v\n", arrayForSort)
	fmt.Printf("Отсортированный массив: %v\n", bubbleSort(arrayForSort))
}
