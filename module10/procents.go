package main

import (
	"fmt"
	"math"
)

func main() {
	var sum, procent, years int
	fmt.Println("Введите сумму которую вы хотите положить в банк")
	fmt.Scan(&sum)
	fmt.Println("Введите ежемесячный процент капитализации")
	fmt.Scan(&procent)
	fmt.Println("Введите количество лет")
	fmt.Scan(&years)
	countOfMonth := years * 12
	result := float64(sum)
	toBank := 0.0
	for numMonth := 0; numMonth < countOfMonth; numMonth++ {
		divs := result * (float64(procent) / 100.0)
		minDivs := math.Floor(divs*100) / 100
		toBank += divs - minDivs
		result += minDivs
	}
	fmt.Println(result, toBank)
}
