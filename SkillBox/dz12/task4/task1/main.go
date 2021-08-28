package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("*** Урок №6 пакет ioutil | урок №4 ***")

	var userString string
	var line int
	var endString = "выход"
	var fileName = "dz12_task4.txt"
	var b bytes.Buffer

	for {
		fmt.Printf("Введите строку [для выхода введите \"%s\"]: ", endString)
		userString, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		userString = strings.Trim(userString, "\n")

		if userString == endString {
			break
		}

		line++
		date := time.Now()
		b.WriteString(fmt.Sprintf("%d: %s %s\n", line, date.Format("2006-01-02 15:04:05"), userString))
	}
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
}
