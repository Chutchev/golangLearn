package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	for {
		fmt.Println("Введите число: ")
		fmt.Scan(&str)
		if strings.ToLower(str) == "стоп" {
			break
		}
		number, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Произошла ошибкак парсинга")
			continue
		}
		fmt.Printf("Ввод: %d\n", number)
		fs := square(number)
		ss := twice(fs)
		fmt.Printf("Произведение: %d\n", <-ss)

	}
}

func square(number int) chan int {
	result := make(chan int)
	go func() {
		result <- (number * number)
	}()
	return result
}

func twice(numberChan chan int) chan int {
	a := <-numberChan
	fmt.Printf("Квадрат: %d\n", a)
	result := make(chan int)
	go func() {
		result <- 2 * a
	}()
	return result
}
