package main

import (
	"errors"
	"fmt"
)

func main() {
	var matrixFirst [3][5]int
	var maxtrixSecond [5][4]int

	fmt.Println("Заполнение первой матрицы: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("Введите %d строку: \n", i+1)
		for j := 0; j < 5; j++ {
			fmt.Printf("matrixFirst[%d][%d] = ", i+1, j+1)
			fmt.Scan(&matrixFirst[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Заполнение второй матрицы: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("Введите %d строку: \n", i+1)
		for j := 0; j < 4; j++ {
			fmt.Printf("matrixFirst[%d][%d] = ", i+1, j+1)
			fmt.Scan(&maxtrixSecond[i][j])
		}
		fmt.Println()
	}
	result, err := multipleMatrix(matrixFirst, maxtrixSecond)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func multipleMatrix(first [3][5]int, second [5][4]int) ([3][4]int, error) {
	var result [3][4]int
	if len(first[0]) != len(second) {
		return result, errors.New("Невозможно произвести умножение матриц!!!")
	}

	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second[0]); j++ {
			for k := 0; k < len(second); k++ {
				result[i][j] += first[i][k] * second[k][j]
			}
		}
	}
	return result, nil
}
