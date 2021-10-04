package main

import (
	"flag"
	"fmt"
)

func findSubstr(str *string, substr *string) bool {
	substrRunes := []rune(*substr)

	n := 0
	for _, letter := range *str {
		if substrRunes[n] != letter {
			n = 0
			continue
		}

		n++

		if n == len(substrRunes) {
			return true
		}

	}

	return false
}

func main() {
	var (
		str    string
		substr string
	)

	flag.StringVar(&str, "str", "", "set str")
	flag.StringVar(&substr, "substr", "", "set substr")
	flag.Parse()

	fmt.Println(findSubstr(&str, &substr))
}
