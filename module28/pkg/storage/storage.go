package storage

import (
	"fmt"
	"module28/pkg/student"
)

type StudentStorage map[string]*student.Student

func (ss StudentStorage) Put(student *student.Student) (int, error) {
	ss[student.Name] = student

	if ss[student.Name] == nil {
		return -1, fmt.Errorf("Не удалось добавить студента в мапу")
	}
	return 1, nil
}

func (ss StudentStorage) Get(student *student.Student) (*student.Student, error) {
	if ss[student.Name] == nil {
		return nil, fmt.Errorf("Не найден студент с имененем %v", student.Name)
	}
	return ss[student.Name], nil
}

func NewStudentStorage() StudentStorage {
	return make(map[string]*student.Student)
}

func (ss StudentStorage) PrintAllStudents() {
	for _, student := range ss {
		fmt.Println(student.GetInfo())
	}
}
