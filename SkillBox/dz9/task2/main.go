package main

import (
	"fmt"
	"math"
)

func main() {
	var number1 int16
	var number2 int16
	var resType string

	fmt.Print("Введите первое число: ")
	_, _ = fmt.Scan(&number1)
	fmt.Print("Введите второе число: ")
	_, _ = fmt.Scan(&number2)

	result := int32(number1) * int32(number2)

	if result > 0 {
		switch {
		case result <= math.MaxUint8:
			resType = "uint8"
		case result <= math.MaxUint16:
			resType = "uint16"
		default:
			resType = "uint32"
		}
	} else {
		switch {
		case result >= math.MaxInt8:
			resType = "int8"
		case result <= math.MaxInt16:
			resType = "int16"
		default:
			resType = "int32"
		}
	}

	fmt.Printf("Минимальный тип для результата умножения %d * %d = %d: %s\n", number1, number2, result, resType)
}
