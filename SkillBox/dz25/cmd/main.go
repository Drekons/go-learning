package main

import (
	"bufio"
	"dz25/pkg/storage"
	student "dz25/pkg/student"
	"fmt"
	"os"
)

func main() {
	fmt.Println("**** Список студентов ****")
	studentStorage := storage.NewStudentStorage()

	defer func() {
		fmt.Printf("Студенты из хранилища:\n")
		for _, s := range studentStorage {
			fmt.Printf("%s %d %d\n", s.GetName(), s.GetAge(), s.GetGrade())
		}
	}()

	fmt.Println("Водите данные о студентах через пробел в формате Имя Возраст Курс")
	fmt.Println("Для завершения ввода списка студентов нажмите ctrl + D")
	myScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Введите данные студента [Иван 21 3]: ")

		myScanner.Scan()
		studentString := myScanner.Text()
		if studentString == "" {
			fmt.Println()
			return
		}

		studentStorage.Put(student.NewStudent(studentString))
	}
}
