package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8, 9}
	// Первый вариант
	c := appendOther(a, b)
	fmt.Println("Первый вариант")
	fmt.Println(c)
	// Второй вариант

	d := append(a, b...)
	fmt.Println("Второй вариант")
	fmt.Println(d)
}

func appendOther(a, b []int) [9]int {
	var c [9]int
	for i, value := range a {
		c[i] = value
	}
	lenA := len(a)
	for i, value := range b {
		c[lenA+i] = value
	}
	return c
}
