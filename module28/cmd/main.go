package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"module28/pkg/storage"
	"module28/pkg/student"
	"os"
	"strconv"
	"strings"
)

func main() {
	ss := storage.NewStudentStorage()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Введите Имя, возраст и оценку студента\n")
		stroke, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Заканчиваем набор студентов в группу!")
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
		student := student.NewStudent(name, age, grade)
		_, err = ss.Put(student)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("__________________")
	fmt.Println("Студенты: ")
	ss.PrintAllStudents()

}
