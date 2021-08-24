package main

import "fmt"

func main() {
	var day string

	fmt.Print("Введите день недели [пн, вт, ср, чт, пт]: ")
	_, _ = fmt.Scan(&day)

	switch day {
	case "пн":
		fmt.Println("вторник")
		fallthrough
	case "вт":
		fmt.Println("среда")
		fallthrough
	case "ср":
		fmt.Println("четверг")
		fallthrough
	case "чт":
		fmt.Println("пятница")
		fallthrough
	case "пт":
	default:
		fmt.Println("День недели не найден!")
	}
}
