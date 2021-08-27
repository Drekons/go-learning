package main

import (
	"fmt"
	"math"
)

func main() {
	var number1 int32
	var number2 int32
	var resType string

	fmt.Print("Введите первое число: ")
	_, _ = fmt.Scan(&number1)
	fmt.Print("Введите второе число: ")
	_, _ = fmt.Scan(&number2)

	result := int64(number1) * int64(number2)

	if result > 0 {
		switch {
		case result <= math.MaxUint8:
			resType = "uint8"
		case result <= math.MaxUint16:
			resType = "uint16"
		case result <= math.MaxUint32:
			resType = "uint32"
		default:
			resType = "uint64"
		}
	} else {
		switch {
		case result*-1 <= math.MaxInt8:
			resType = "int8"
		case result*-1 <= math.MaxInt16:
			resType = "int16"
		case result*-1 <= math.MaxInt32:
			resType = "int32"
		default:
			resType = "int64"
		}
	}

	fmt.Printf("Минимальный тип для результата умножения %d * %d = %d: %s\n", number1, number2, result, resType)
}
