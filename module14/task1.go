package main

import "fmt"

func isOdd(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	var n int
	fmt.Println("Введите число: ")
	fmt.Scan(&n)
	fmt.Print(isOdd(3))
}
