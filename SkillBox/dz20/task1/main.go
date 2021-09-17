package main

import "fmt"

func oddOrEven(arr []int) (even []int, odd []int) {
	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return
}

func main() {
	fmt.Println("*** Задание 1. Чётные и нечётные ***")

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("Исходный массив %v\n", arr)
	even, odd := oddOrEven(arr)

	fmt.Printf("Чётный массив %v\n", even)
	fmt.Printf("Нечётный массив %v\n", odd)
}
