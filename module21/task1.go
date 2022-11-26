package main

import "fmt"

func solveProblem(x int16, y uint8, z float32) float32 {
	return 2*float32(x) + float32(y*y) - 3/z
}

func main() {
	S := solveProblem(1, 2, 4)
	fmt.Println(S)
}
