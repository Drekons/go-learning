package student

import (
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

func NewStudent(studentString string) *Student {
	args := strings.Split(studentString, " ")
	s := &Student{}
	s.PutName(args[0])
	s.PutAgeString(args[1])
	s.PutGradeString(args[2])
	return s
}
