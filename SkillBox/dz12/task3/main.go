package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("*** Урок №4 уровни доступа ***")

	fileName := "dz12_task1.txt"
	file := openFile(fileName)
	if err := os.Chmod(fileName, 0444); err != nil {
		panic(err)
	}
	defer closeFile(file)

	_, err := file.WriteString("test")
	if err != nil {
		fmt.Printf("В файл \"%s\" нет возможности записи\n", fileName)
	} else {
		fmt.Printf("Запись в файл \"%s\" прошла успешно\n", fileName)
	}
}

func openFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	return file
}

func closeFile(file *os.File) {
	_ = file.Close()
}
