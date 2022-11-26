package main

import "fmt"

func someFunc(a, b int, A func(int, int) int) {
	defer A(a, b)
	fmt.Println("Вызывается функция someFunc")
}

func main() {
	someFunc(1, 2, func(a, b int) int { fmt.Printf("%d + %d = %d\n", a, b, a+b); return a + b })
	someFunc(3, 5, func(a, b int) int { fmt.Printf("%d - %d = %d\n", a, b, a-b); return a - b })
	someFunc(6, 8, func(a, b int) int { fmt.Printf("%d * %d = %d\n", a, b, a*b); return a * b })
}
