package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s *Student) PutName(name string) {
	s.name = name
}
func (s *Student) PutAge(age int) {
	s.age = age
}
func (s *Student) PutGrade(grade int) {
	s.grade = grade
}
func (s *Student) GetName() string {
	return s.name
}
func (s *Student) GetAge() int {
	return s.age
}
func (s *Student) GetGrade() int {
	return s.grade
}
func (s *Student) PutAgeString(age string) {
	ageInt, _ := strconv.Atoi(age)
	s.PutAge(ageInt)
}
func (s *Student) PutGradeString(grade string) {
	gradeInt, _ := strconv.Atoi(grade)
	s.PutGrade(gradeInt)
}
func (s *Student) newStudent(studentString string) *Student {
	args := strings.Split(studentString, " ")
	s.PutName(args[0])
	s.PutAgeString(args[1])
	s.PutGradeString(args[2])
	return s
}

func main() {
	students := make(map[string]*Student)

	defer func() {
		fmt.Printf("Студенты из хранилища:\n")
		for _, student := range students {
			fmt.Printf("%s %d %d\n", student.GetName(), student.GetAge(), student.GetGrade())
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

		student := new(Student)
		student.newStudent(studentString)
		students[student.name] = student
	}
}
