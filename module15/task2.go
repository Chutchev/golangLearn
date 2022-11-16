package main

import (
	"fmt"
	"strconv"
	"strings"
)

const SIZE = 10

func reverseArray(array []int) []int {
	var reversedArray [SIZE]int
	for i := 10; i > 0; i-- {
		reversedArray[10-i] = array[i-1]
	}

	return reversedArray[:]
}

func main() {

	var numbers [SIZE]int
	var number string
	fmt.Printf("Вводите числа. Для прекращения ввода введите exit. Максимальное количество чисел: %d\n", SIZE)

	for i := 0; i < SIZE; i++ {

		fmt.Scan(&number)
		number = strings.ToLower(number)
		if number == "exit" {
			break
		}
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			fmt.Printf("Нельзя конвертировать %q в Int\n", number)
			panic(err)
		}
		numbers[i] = numberInt
	}
	reversedArray := reverseArray(numbers[:])
	fmt.Println(reversedArray)
}
