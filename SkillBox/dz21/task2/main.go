package main

import (
	"fmt"
	"math/rand"
	"time"
)

func modify(array []int, modifyFunc func([]int)) {
	modifyFunc(array)
}

func main() {
	fmt.Println("Задание 2. Анонимные функции ")
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)

	for k := range arr {
		arr[k] = rand.Intn(k + 10)
	}

	fmt.Printf("Исходный массив %+v\n", arr)

	modify(arr, func(array []int) {
		for i := 0; i < len(array)-1; i++ {
			for j := 0; j < len(array)-i-1; j++ {
				if array[j] > array[j+1] {
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}
	})

	fmt.Printf("Сортировка пузырьком массив %+v\n", arr)

	modify(arr, func(array []int) {
		for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
			array[i], array[j] = array[j], array[i]
		}
	})

	fmt.Printf("Перевёрнутый массив %+v\n", arr)

}
