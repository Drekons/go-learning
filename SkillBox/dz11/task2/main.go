package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("*** Вывод чисел содержащихся в строке ***")

	var str = "a10 10 20b 20 30c30 30 dd"
	var result = ""

	fmt.Printf("Исходная строка: %s\n", str)

	str = strings.Trim(str, " ")
	for len(str) > 0 {
		var part string
		spaceIndex := strings.Index(str, " ")
		if spaceIndex > 0 {
			part = str[:spaceIndex]
			str = str[spaceIndex:]
		} else {
			part = str
			str = ""
		}

		str = strings.Trim(str, " ")

		_, err := strconv.Atoi(part)
		if err == nil {
			result += " " + part
		}
	}

	result = strings.Trim(result, " ")

	fmt.Printf("Части строки которые можно привест к целому числу: %s\n", result)
}
