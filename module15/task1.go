package main

import "fmt"

func main() {
	var numbers [10]int
	var odd, even int
	fmt.Println("Введите 10 чисел: ")
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}
	for _, number := range numbers {
		if number%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Printf("Четных чисел: %d, Нечетных: %d\n", even, odd)
}
