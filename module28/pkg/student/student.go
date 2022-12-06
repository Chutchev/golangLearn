package student

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name string, age, grade int) *Student {
	return &Student{Name: name, Age: age, Grade: grade}
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("Name: %v, Age: %d, Grade: %d", s.Name, s.Age, s.Grade)
}
