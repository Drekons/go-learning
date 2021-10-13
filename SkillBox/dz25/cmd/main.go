package main

import (
	"bufio"
	"dz25/pkg/storage"
	student "dz25/pkg/student"
	"fmt"
	"os"
)

func main() {
	studentStorage := storage.NewStudentStorage()

	defer func() {
		fmt.Printf("Студенты из хранилища:\n")
		for _, s := range studentStorage {
			fmt.Printf("%s %d %d\n", s.GetName(), s.GetAge(), s.GetGrade())
		}
	}()

	myScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Введите данные студента: ")

		myScanner.Scan()
		studentString := myScanner.Text()
		if studentString == "" {
			fmt.Println()
			return
		}

		studentStorage.Put(student.NewStudent(studentString))
	}
}
