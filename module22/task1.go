// Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число. +
// Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
// При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	rand.Seed(time.Now().Unix())
	var array [n]int
	var num int
	fillArray(&array)
	fmt.Println(array)
	fmt.Println("Введите число: ")
	fmt.Scan(&num)
	count := countAfterNum(array, num)
	fmt.Printf("Количество чисел после заданного: %d\n", count)
}

func fillArray(a *[n]int) {
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n * n)
	}
}

func countAfterNum(a [n]int, num int) (count int) {
	count = 0
	for i, nu := range a {
		if nu == num {
			count = len(a) - (i + 1)
			break
		}
	}
	return
}
