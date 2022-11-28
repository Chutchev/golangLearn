// Заполните упорядоченный массив из 12 элементов и введите число.
// Необходимо реализовать поиск первого вхождения заданного числа в массив.
// Сложность алгоритма должна быть минимальная.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 12

func main() {
	rand.Seed(time.Now().Unix())
	var array [n]int
	fillArray(&array)
	fmt.Println(array)
	fmt.Println("Введите число которое надо найти индекс: ")
	var n int
	fmt.Scan(&n)
	fmt.Printf("Индекс: %d", indexOf(array, n))
}

func fillArray(a *[n]int) {
	for i := 0; i < n; i++ {
		a[i] = 10*i + rand.Intn(n)
	}
}

func indexOf(array [n]int, num int) int {
	left := 0
	high := len(array) - 1
	for left <= high {
		mid := left + (high-left)/2
		if num < array[mid] {
			high = mid - 1
		} else if num > array[mid] {
			left = mid + 1
		} else if mid == 0 || num != array[mid-1] {
			return mid
		} else {
			high = mid - 1
		}
	}

	return -1
}
