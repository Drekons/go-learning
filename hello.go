package main

import "fmt"

func main() {
	arr := []string{"test", "string", "world"}

	fmt.Println("Hello, world!", in_array("test", arr))
	fmt.Println("Hello, world!", in_array("asdad", arr))
	fmt.Println("Hello, world!", in_array("world", arr))

}

func in_array(element string, array []string) bool {
	for i := range array {
		if array[i] == element {
			return true
		}
	}

	return false
}
