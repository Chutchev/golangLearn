package main

import "fmt"

func multipleN(n int) (multiple int) {
	multiple = 2 * n
	return
}

func sumN(n int) (sum int) {
	sum = 2 + n
	return
}

func main() {
	var n int
	fmt.Println("Введите число: ")
	fmt.Scan(&n)

	fmt.Println(multipleN(n))
	fmt.Println(sumN(n))
}
