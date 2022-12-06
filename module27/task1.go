package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func newStudent(name string, age, grade int) Student {
	return Student{Name: name, Age: age, Grade: grade}
}

func main() {
	studentMap := make(map[string]*Student)
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Printf("Введите Имя, возраст и оценку студента\n")
		stroke, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		stroke = strings.Trim(stroke, "\n")
		info := strings.Split(stroke, " ")
		name := info[0]
		age, errAge := strconv.Atoi(info[1])
		grade, errGrade := strconv.Atoi(info[2])
		if errGrade != nil || errAge != nil {
			fmt.Println(errAge, errGrade)
			fmt.Printf("Ошибка в оценке=%v либо в возрасте=%v\n", info[1], info[2])
			continue
		}
		student := newStudent(name, age, grade)
		_, err = put(studentMap, &student)
	}
	for _, student := range studentMap {
		fmt.Printf("Student name: %v, Age: %v, Grade: %v\n", student.Name, student.Age, student.Grade)
	}
}

func put(studentMap map[string]*Student, student *Student) (int, error) {
	studentMap[student.Name] = student

	if studentMap[student.Name] == nil {
		return -1, fmt.Errorf("Не удалось добавить студента в мапу")
	}
	return 1, nil
}

func get(studentMap map[string]*Student, student *Student) (*Student, error) {
	if studentMap[student.Name] == nil {
		return nil, fmt.Errorf("Не найден студент с имененем %v", student.Name)
	}
	return studentMap[student.Name], nil
}
