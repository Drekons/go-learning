package main

import "fmt"

func square(x *int) {
	if x != nil {
		*x = (*x) * (*x)
	}
}

func decideTwo(x *int) {
	if x != nil {
		*x = (*x) / 2
	}
}

func modify(x int, secondFunc func(*int), firstFunc func(*int)) int {
	firstFunc(&x)
	secondFunc(&x)
	return x
}

func main() {
	fmt.Println(modify(10, decideTwo, square))
}
