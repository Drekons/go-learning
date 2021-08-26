package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	var eps int

	fmt.Print("Введите x: ")
	_, _ = fmt.Scan(&x)

	fmt.Print("Введите точность: ")
	_, _ = fmt.Scan(&eps)

	var epsilon = 1. / math.Pow(10., float64(eps))

	res := 1.
	a := 1.

	for n := 1; a > epsilon; n++ {
		a *= float64(x) / float64(n)
		res += a
	}

	fmt.Printf("%v\n", res)
}
