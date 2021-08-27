package main

import (
	"fmt"
	"strings"
)

func getCountUpCase(str string) int {
	var count int

	str = strings.Trim(str, " ")

	for len(str) > 0 {
		first := str[:1]
		if first == strings.ToUpper(first) {
			count++
		}

		nextSpace := strings.Index(str, " ")
		if nextSpace <= 0 {
			break
		}

		str = str[nextSpace:]
		str = strings.Trim(str, " ")
	}

	return count
}

func test(str string, count int) {
	var result = getCountUpCase(str)
	fmt.Printf("Test case \"%s\"\n", str)
	fmt.Printf("В строке %d слов начинающихся с заглавной буквы", result)

	if result == count {
		fmt.Println(" -- OK")
	} else {
		fmt.Println(" -- ERROR")
	}
}

func main() {
	fmt.Println("*** Определение количества слов начинающихся с большой буквы ***")

	test("Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software", 5)
	fmt.Println("-----------")
	test("Go is an Open source programming LangUage that makes it Easy to build simPle, reliable, and efficient Software", 5)
	fmt.Println("-----------")
	test("", 0)
	fmt.Println("-----------")
	test(" ", 0)
	fmt.Println("-----------")
	test("A", 1)
	fmt.Println("-----------")
	test("a", 0)
	fmt.Println("-----------")
}
