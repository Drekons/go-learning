package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("*** Урок №6 пакет ioutil | Урок №3 ***")

	fileName := "dz12_task1.txt"
	readBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", readBytes)
}
