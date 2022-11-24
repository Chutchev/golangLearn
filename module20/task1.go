package main

import "fmt"

const SIZE = 3

func findDeterminant(a [SIZE][SIZE]int) int {
	first := a[1][1]*a[2][2] - a[1][2]*a[2][1]
	second := a[1][0]*a[2][2] - a[1][2]*a[2][0]
	third := a[1][0]*a[2][1] - a[1][1]*a[2][0]
	return a[0][0]*first - a[0][1]*second + a[0][2]*third
}

func main() {
	var a [SIZE][SIZE]int
	for i := 0; i < SIZE; i++ {
		fmt.Printf("Введите %d строку: \n", i+1)
		for j := 0; j < SIZE; j++ {
			fmt.Printf("a[%d][%d] = ", i+1, j+1)
			fmt.Scan(&a[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Исходная матрица")
	for i := 0; i < SIZE; i++ {
		fmt.Printf(" %d | %d | %d\n", a[i][0], a[i][1], a[i][2])
	}
	fmt.Printf("\nДетерминант: %d\n", findDeterminant(a))
}
