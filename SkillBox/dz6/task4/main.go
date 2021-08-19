package main

import "fmt"

func main() {
	fmt.Println("*** Предыдущий урок на if ***")

	var first int
	var second int
	var third int

	firstLimit := 10
	secondLimit := 100
	thirdLimit := 1000

	for first != firstLimit || second != secondLimit || third != thirdLimit {
		if first < firstLimit {
			first++
		}
		if second < secondLimit {
			second++
		}
		if third < thirdLimit {
			third++
		}

		fmt.Println("Переменная 1:", first)
		fmt.Println("Переменная 2:", second)
		fmt.Println("Переменная 3:", third)
	}
}
