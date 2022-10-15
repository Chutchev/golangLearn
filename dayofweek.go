package main

import "fmt"

func main() {
	fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
	var dayOfWeek string
	_, _ = fmt.Scan(&dayOfWeek)
	switch dayOfWeek {
	case "пн":
		fmt.Println("Понедельник")
		fallthrough
	case "вт":
		fmt.Println("Вторник")
		fallthrough
	case "ср":
		fmt.Println("Среда")
		fallthrough
	case "чт":
		fmt.Println("Четверг")
		fallthrough
	case "пт":
		fmt.Println("Пятница")
	default:
		fmt.Println("Я не знаю такого дня")
	}
}
