package main

import (
	"fmt"
	"math"
)

// exponent = 1 + x^k/x!
func main() {
	fmt.Println("Введите x для которого необходимо рассчитать ex")
	var x, n float64
	fmt.Scan(&x)
	fmt.Println("Введите n точность после запятой")
	fmt.Scan(&n)
	epsilon := 1.0 / math.Pow(10.0, float64(n))
	result := 1.0
	prevRes := 1.0
	fact := 1
	k := 1
	for {
		if k != 1 {
			fact *= k
		}
		result += math.Pow(x, float64(k)) / float64(fact)
		if math.Abs(float64(result))-float64(prevRes) < epsilon {
			fmt.Print(result)
			break
		}
		k++
		prevRes = result
	}
}
