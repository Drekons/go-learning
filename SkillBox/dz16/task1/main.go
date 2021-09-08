package main

import "fmt"

const array1Len = 4
const array2Len = 5

func mixArray(array1 [array1Len]int, array2 [array2Len]int) (resultArray [array1Len + array2Len]int) {
	resultLen := array1Len + array2Len

	for i := 0; i < resultLen; i++ {
		if i >= array1Len {
			resultArray[i] = array2[i-array1Len]
		} else {
			resultArray[i] = array1[i]
		}
	}

	return
}

func main() {
	fmt.Println("Задание 1. Слияние отсортированных массивов")

	array1 := [array1Len]int{1, 2, 3, 4}
	array2 := [array2Len]int{5, 6, 7, 8, 9}

	fmt.Printf("Результат слияния массива %v и %v = %v\n", array1, array2, mixArray(array1, array2))
}
