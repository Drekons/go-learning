package main

import (
	"fmt"
	"math"
)

func main() {
	var countUnit8 = 0
	var countUnit16 = 0

	for i := 0; i <= math.MaxUint32; i++ {
		if i > math.MaxUint8 {
			countUnit8++
		}
		if i > math.MaxUint16 {
			countUnit16++
		}
	}

	fmt.Printf("Переполнений для unti8 %d, переполнений unit16 %d\n", countUnit8, countUnit16)
}
