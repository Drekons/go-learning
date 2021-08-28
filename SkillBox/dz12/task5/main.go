package main

import (
	"fmt"
	"strings"
)

func main() {
	var count int
	var stack []string

	fmt.Print("Введите кол-во открывающихся скобок: ")
	_, _ = fmt.Scan(&count)

	if count < 0 {
		fmt.Printf("%v\n", stack)
		return
	}

	makeParentheses(count, count, "", &stack)

	fmt.Printf("%v\n", stack)
}

func makeParentheses(left, right int, temp string, result *[]string) {
	if left == 0 {
		*result = append(*result, temp+strings.Repeat(")", right))
		return
	}
	if right > left {
		makeParentheses(left, right-1, temp+")", result)
	}
	makeParentheses(left-1, right, temp+"(", result)
}
