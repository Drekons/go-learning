package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	stopWord := "стоп"

	fmt.Println("Введите числа для расчёта квадрата и его произведения")
	fmt.Printf("Для завершения программы введите \"%s\"\n", stopWord)
	for {
		wg.Wait()
		fmt.Printf("Введите число или слово \"%s\": ", stopWord)
		var input string
		_, _ = fmt.Scan(&input)

		if input == stopWord {
			fmt.Println("Программа завершена")
			return
		}

		inputInt, er := strconv.Atoi(input)
		if er != nil {
			fmt.Println("Введено не корректное значение")
			continue
		}
		wg.Add(1)
		go modify(inputInt, &wg)
	}
}

func modify(i int, wg *sync.WaitGroup) {
	firstChannel := toSquare(i)
	secondChannel := toMultipleTwo(firstChannel)

	fmt.Printf("Получено значение %d | квадрат %d | произведение от квадрата %d\n", i, <-secondChannel, <-secondChannel)
	wg.Done()
}

func toSquare(i int) chan int {
	channel := make(chan int)
	go func() {
		channel <- i * i
	}()
	return channel
}

func toMultipleTwo(channel chan int) chan int {
	nextChannel := make(chan int)
	go func() {
		square := <-channel
		nextChannel <- square
		nextChannel <- square * 2
	}()

	return nextChannel
}
