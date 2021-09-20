package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertSort(arr []int) {
	for k, v := range arr {
		current := v
		newKey := k
		for j := k - 1; j > -1; j-- {
			if arr[j] < current {
				break
			}
			arr[j+1] = arr[j]
			newKey = j
		}
		arr[newKey] = current
	}
}

func main() {
	fmt.Println("*** Задание 1. Сортировка вставками ***")

	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)

	for k := range arr {
		arr[k] = rand.Intn(k + 10)
	}

	fmt.Printf("Исходный массив %+v\n", arr)
	insertSort(arr)
	fmt.Printf("Отсортированный массив %+v\n", arr)
}
