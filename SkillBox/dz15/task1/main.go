package main

import "fmt"

func main() {
	fmt.Println("*** Задание 1. Подсчёт чётных и нечётных чисел в массиве ***")

	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Введите число №%d: ", i+1)
		_, _ = fmt.Scan(&numbers[i])
	}

	var even, odd int

	for _, number := range numbers {
		if number%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Printf("чётных — %d, нечётных — %d\n", even, odd)
}
