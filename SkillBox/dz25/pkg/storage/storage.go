package storage

import "dz25/pkg/student"

type StudentStorage map[string]*student.Student

func NewStudentStorage() StudentStorage {
	return make(map[string]*student.Student)
}

func (ss StudentStorage) Put(s *student.Student) {
	ss[s.GetName()] = s
}
