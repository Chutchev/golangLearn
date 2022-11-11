package main

import "fmt"

func sumOfEven(a, b int) {
	max := a
	min := b
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	sum := 0
	for i := min; i < max; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

func main() {
	var a, b int
	fmt.Println("Введите два числа")
	fmt.Scan(&a, &b)
	sumOfEven(a, b)
}
