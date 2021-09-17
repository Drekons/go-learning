package main

import "fmt"

func parseTest(sentences []string, chars []rune) {

}

func main() {
	fmt.Println("*** Задание 2. Поиск символов в нескольких строках ***")

	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}

	fmt.Printf("Массив предложений %v\n", sentences)
	fmt.Printf("Массив символов %v\n", chars)
}
