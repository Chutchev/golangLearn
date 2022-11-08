package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "a10 10 20b 20 30c30 30 dd"
	fmt.Println("В строке содержатся числа в десятичном формате:")
	for len(a) > 0 {
		spaceIndex := strings.Index(a, " ")
		if spaceIndex == -1 {
			break
		}
		word := a[:spaceIndex]

		number, err := strconv.Atoi(word)
		if err == nil {
			fmt.Print(number, " ")
		}
		a = a[spaceIndex+1:]
		// fmt.Println("a: ", a)
	}
}
