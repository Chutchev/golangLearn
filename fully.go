package main

import (
	"fmt"
)

// const (
// 	maxInt8  = math.MaxInt8
// 	maxInt16 = math.MaxInt16
// 	maxInt32 = math.MaxInt32
// )

func main() {
	summa8 := 0
	summa16 := 0
	for i := maxInt8; i < maxInt32; i++ {
		if int32(i) > int32(maxInt16) {
			summa16++
		}
		summa8++
	}
	fmt.Printf("Переполнений int8: %d\n", summa8)
	fmt.Printf("Переполнений int16: %d\n", summa16)
}
