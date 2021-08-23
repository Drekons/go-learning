package main

import (
	"fmt"
)

func main() {
	fmt.Println("*** Вывод елочки ***")

	var length int

	fmt.Print("Введите колво строк для ёлочки: ")
	_, _ = fmt.Scan(&length)

	canvasSize := length*2 + 1
	for i := 1; i <= length; i++ {
		count := i + i - 1
		skip := (canvasSize - count) / 2

		for j := 0; j < canvasSize; j++ {
			if j < skip || j > skip+count-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Print("\n")
	}
}
