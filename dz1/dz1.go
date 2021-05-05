package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("String: %s unpack: %s\n", "a4bc2d5e", unpackStr("a4bc2d5e"))
	fmt.Printf("String: %s unpack: %s\n", "abcd", unpackStr("abcd"))
	fmt.Printf("String: %s unpack: %s\n", "45", unpackStr("45"))
	fmt.Printf("String: %s unpack: %s\n", `qwe\4\5`, unpackStr(`qwe\4\5`))
	fmt.Printf("String: %s unpack: %s\n", `qwe\45`, unpackStr(`qwe\45`))
	fmt.Printf("String: %s unpack: %s\n", `qwe\\5`, unpackStr(`qwe\\5`))
}

func unpackStr(str string) string {
	var newSting string = ""
	var symbol string = ""
	var rSymbol string = ""
	var slash bool = false
	var first bool = false

	for i := 0; i < len(str); i++ {
		symbol = string(byte(str[i]))
		matched, err := regexp.MatchString(`\d`, symbol)
		if symbol == `\` && !slash {
			slash = true
			continue
		}
		if matched && err == nil && !slash {
			if len(rSymbol) > 0 {
				intSymbol, err := strconv.Atoi(symbol)
				if err == nil {
					newSting += reapetedString(rSymbol, intSymbol)
				}
				first = false
			}
		} else {
			if first {
				newSting += rSymbol
			}
			first = true
			rSymbol = symbol
		}
		slash = false
	}

	if first && len(rSymbol) > 0 {
		newSting += rSymbol
	}

	return newSting
}

func reapetedString(str string, count int) string {
	var newString string = ""

	for i := 0; i < count; i++ {
		newString += str
	}

	return newString
}
