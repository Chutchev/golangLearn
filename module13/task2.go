package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b int
	a = 10
	b = 20
	fmt.Printf("Старые значения: %d, %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("Новые значения: %d, %d\n", a, b)
}
