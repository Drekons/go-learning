package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func getS(x int16, y uint8, z float32) float32 {
	S := 2.*float32(x) + float32(math.Pow(float64(y), 2.)) - 3./z
	return S
}

func main() {
	fmt.Println("*** Задание 1. Расчёт по формуле ***")

	rand.Seed(time.Now().UnixNano())
	x := int16(rand.Intn(math.MaxInt16 - 1))
	y := uint8(rand.Intn(math.MaxUint8 - 1))
	z := rand.Float32()

	fmt.Printf("2 * %v + %v ^ 2 - 3 / %v = %v\n", x, y, z, getS(x, y, z))
}
