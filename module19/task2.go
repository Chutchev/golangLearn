package main

import "fmt"

func main() {
	a := [6]int{6, 4, 2, 1, 3, 5}
	b := bubbleSort(a)
	fmt.Println(b)
}

func bubbleSort(a [6]int) [6]int {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
