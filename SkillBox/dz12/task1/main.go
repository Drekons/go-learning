package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("*** Урок №2 Работа с файлами ***")

	var userString string
	var line int
	var endString = "выход"

	file := createFile("dz12_task1.txt")
	defer closeFile(file)

	for {
		fmt.Printf("Введите строку [для выхода введите \"%s\"]: ", endString)
		userString, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		userString = strings.Trim(userString, "\n")

		if userString == endString {
			break
		}

		line++
		date := time.Now()
		writingStr := fmt.Sprintf("%d: %s %s\n", line, date.Format("2006-01-02 15:04:05"), userString)
		_, err := file.WriteString(writingStr)
		if err != nil {
			panic(err)
		}
	}
}

func createFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	return file
}

func closeFile(file *os.File) {
	_ = file.Close()
}
