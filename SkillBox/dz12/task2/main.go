package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("*** Урок №3 интерфейс io.Reader ***")

	file := openFile("dz12_task1.txt")
	defer closeFile(file)

	buf := make([]byte, getFileSize(file))
	if _, err := io.ReadFull(file, buf); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", buf)
}

func openFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	return file
}

func getFileSize(file *os.File) int64 {
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	return stat.Size()
}

func closeFile(file *os.File) {
	_ = file.Close()
}
