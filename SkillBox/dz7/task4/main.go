package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("*** Сколько нужно купить билетов, для получения счастливого ***")

	var maxCount int
	var steps int
	var length = 6

	for i := 100000; i <= 999999; i++ {
		steps++

		var left int
		var right int
		for j := 1; j < 4; j++ {
			left += i / int(math.Pow10(length-j)) % 10
			right += i % int(math.Pow10(j)) / int(math.Pow10(j-1))
		}

		if left == right {
			if steps > maxCount {
				maxCount = steps
			}
			steps = 0
		}
	}

	fmt.Println("Нужно купить", maxCount, "билет")
}
