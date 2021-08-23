package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("*** Зеркальные билеты ***")

	var length = 6
	var count int

numbersFor:
	for i := 100000; i <= 999999; i++ {
		for j := 1; j < 4; j++ {
			left := i / int(math.Pow10(length-j)) % 10
			right := i % int(math.Pow10(j)) / int(math.Pow10(j-1))
			if left != right {
				continue numbersFor
			}
		}
		count++
	}

	fmt.Println("Зеркальных билетов:", count)
}
