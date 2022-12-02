package main

import (
	"fmt"
)

const size = 10

func main() {
	odd, even := divideOddEven(1, 3, 10, 14, 7, 19, 22)
	// a := []int{1, 3, 10, 14, 7, 19, 22}
	// odd, even := divideOddEven(a...)
	fmt.Println("---ODD---")
	fmt.Println(odd)
	fmt.Println("---EVEN---")
	fmt.Println(even)
}

func divideOddEven(array ...int) ([]int, []int) {
	var odd, even []int
	for _, value := range array {
		if value%2 == 1 {
			odd = append(odd, value)
		} else {
			even = append(even, value)
		}
	}
	return odd, even
}
