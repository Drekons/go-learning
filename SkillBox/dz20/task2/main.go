package main

import "fmt"

const sentenceCount = 4
const charsCount = 5

func parseTest(sentences [sentenceCount]string, chars [charsCount]rune) (result [sentenceCount][charsCount]int) {
	for i, sentence := range sentences {
		lastWordChars := getLastWordLikeArray(sentence)
		for j, char := range chars {
			result[i][j] = findRuneInArray(char, lastWordChars)
		}
	}

	return
}

func getLastWordLikeArray(sentence string) (lastWordArr []rune) {
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			lastWordArr = []rune{}
			continue
		}

		lastWordArr = append(lastWordArr, rune(sentence[i]))
	}

	return
}

func findRuneInArray(char rune, wordChars []rune) (index int) {
	index--

	for k, wordChar := range wordChars {
		if wordChar == char {
			index = k
			break
		}
	}

	return
}

func main() {
	fmt.Println("*** Задание 2. Поиск символов в нескольких строках ***")

	sentences := [sentenceCount]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [charsCount]rune{'o', 'е', 'i', 'd', 'l'}

	fmt.Printf("Массив предложений %v\n", sentences)
	fmt.Printf("Массив символов %v\n", chars)
	fmt.Printf("Вхождения символов в последнее слово каждого предложения %v\n", parseTest(sentences, chars))
}
